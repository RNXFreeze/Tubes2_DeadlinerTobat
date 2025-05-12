'use client';
import Tree from 'react-d3-tree';
import { useState, useEffect } from 'react';

export default function TreeVisualizer({ target }) {
  const [treeData, setTreeData] = useState(null);
  const [incomingTrees, setIncomingTrees] = useState([]);

  useEffect(() => {
    if (!target) return;

    setIncomingTrees([]); // reset saat target baru

    const baseURL = 'http://localhost:8080/api';
    const es = new EventSource(`${baseURL}/bfs/stream?target=${encodeURIComponent(target)}`);

    es.onmessage = (e) => {
      const newTree = JSON.parse(e.data);
      console.log('SSE-node:', newTree);

      setIncomingTrees(prev => {
        const updated = [...prev, newTree];
        const combined = combineTrees(updated, target);
        setTreeData(combined);
        return updated;
      });
    };

    es.onerror = () => es.close();
    return () => es.close();
  }, [target]);

  return (
    <div style={{ width: '100%', height: '100vh', overflow: 'auto', background: '#fafafa' }}>
      {treeData && (
        <Tree
          data={treeData}
          orientation="vertical"
          collapsible={false}
          translate={{ x: 600, y: 100 }}
          zoom={0.6}
          zoomable={true}
          pathFunc="elbow"
          nodeSize={{ x: 120, y: 100 }}
          separation={{ siblings: 0.6, nonSiblings: 0.8 }}
          renderCustomNodeElement={({ nodeDatum }) => {
            const isLeaf = !nodeDatum.children || nodeDatum.children.length === 0;
            if (nodeDatum.name === '') {
              return <g></g>;
            }
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

function mergeIntoRoot(root, newTree) {
  if (root.name === newTree.name) {
    // gabungkan children dari tree baru ke root
    root.children = [...(root.children || []), ...(newTree.children || [])];

    // root.children = mergeChildren(root.children || [], newTree.children || []);
  } else {
    // cari node yang cocok di bawah root, lalu merge
    insertNodeRecursively(root, newTree);
  }
}

// function mergeChildren(oldChildren, newChildren) {
//   const merged = [...oldChildren];

//   for (const newChild of newChildren) {
//     const existing = merged.find(c => c.name === newChild.name);

//     if (existing) {
//       existing.children = mergeChildren(existing.children || [], newChild.children || []);
//     } else {
//       merged.push(newChild);
//     }
//   }

//   return merged;
// }

function insertNodeRecursively(current, incoming) {
  if (current.name === incoming.name) {
    current.children = [...(current.children || []), ...(incoming.children || [])];
    return true;
  }

  if (!current.children) return false;

  for (const child of current.children) {
    if (insertNodeRecursively(child, incoming)) return true;
  }

  return false;
}

function combineTrees(trees, rootName = "Root") {
  if (!Array.isArray(trees) || trees.length === 0) return null;

  const recipePairs = [];

  for (const tree of trees) {
    if (!tree || !Array.isArray(tree.children)) continue;

    // Bungkus masing-masing tree agar tidak tercampur di visual
    recipePairs.push({
      name: '',
      children: tree.children
    });
  }

  return {
    name: rootName,
    children: recipePairs
  };
}
