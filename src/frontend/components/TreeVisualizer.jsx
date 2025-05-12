'use client'
import Tree from 'react-d3-tree';
import { useState, useEffect } from 'react';

export default function TreeVisualizer({data}) {
  const [treeData, setTreeData] = useState(null);
  const useDummy = false;

  useEffect(() => {
    if(useDummy){
        // Dummy Test
        fetch('/dummy_tree.json')
          .then(res => res.json())
          .then(data => {
            // Start with just root
            const root = JSON.parse(JSON.stringify(data));
            root.children = []; // kosongin anaknya dulu
            setTreeData(root);
    
            // Simulasikan live traversal
            treeLiveVisualization(root, data, setTreeData);
          });
    } else {
        // Ini contoh kalau pake websocket
        const root = { name: '', children: [] };
        setTreeData(root);

        const socket = new WebSocket('ws://localhost:8080/'); // ganti URL sesuai server backend-mu

        socket.onmessage = (event) => {
            const newNode = JSON.parse(event.data);
            updateTreeLive(root, newNode, setTreeData);
        };

        socket.onerror = (err) => {
            console.error('WebSocket error:', err);
        };

        return () => {
            socket.close(); // cleanup kalau component unmount
        };
    }
  }, []);

  return (
    <div style={{ width: '100%', height: '600px' }}>
      {treeData && (
        <Tree
          data={treeData}
          orientation="vertical"
          collapsible={false}
          translate={{ x: 300, y: 50 }}
          pathFunc="elbow"
          shouldRenderLabel={false}
          renderCustomNodeElement={({ nodeDatum }) => {
            const isLeaf = !nodeDatum.children || nodeDatum.children.length === 0;

            return (
                <g>
                <circle r={15} fill={isLeaf ? 'white' : 'purple'} stroke="purple" strokeWidth={2} />
                <text
                    y={-30}
                    x={0}
                    textAnchor="middle" 
                    alignmentBaseline="middle"
                    style={{
                    fontSize: '16px',
                    fontFamily: 'system-ui, sans-serif',
                    fontWeight: 300,
                    fill: 'black',
                    textShadow: 'none'
                    }}
                >
                    {nodeDatum.name}
                </text>
                </g>
            );
          }}
        />
      )}
    </div>
  );
}

function updateTreeLive(currentRoot, incomingNode, setTreeData) {
  const clone = JSON.parse(JSON.stringify(currentRoot));

  function insertNode(current, target) {
    if (current.name === '' || current.name === target.name) {
      current.name = target.name;
      current.children = [];
    }

    if (target.children && target.children.length > 0) {
      for (const child of target.children) {
        let existing = current.children.find(c => c.name === child.name);
        if (!existing) {
          existing = { name: child.name, children: [] };
          current.children.push(existing);
        }
        insertNode(existing, child);
      }
    }
  }

  insertNode(clone, incomingNode);
  setTreeData(clone);
}

async function treeLiveVisualization(liveRoot, fullData, setTreeData) {
  const queue = [{ liveNode: liveRoot, fullNode: fullData }];

  while (queue.length > 0) {
    const { liveNode, fullNode } = queue.shift();

    if (fullNode.children) {
      for (const child of fullNode.children) {
        const newChild = { name: child.name, children: [] };
        liveNode.children.push(newChild);
        setTreeData(JSON.parse(JSON.stringify(liveRoot))); // force re-render
        await new Promise(res => setTimeout(res, 300)); // delay update

        queue.push({ liveNode: newChild, fullNode: child });
      }
    }
  }
}