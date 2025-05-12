'use client';
import Tree from 'react-d3-tree';
import { useState, useEffect } from 'react';

export default function TreeVisualizer({ data }) {
  const [treeData, setTreeData] = useState(null);     // data yang ditampilkan
  const [fullTree, setFullTree] = useState(null);     // data lengkap dari API

  // Jalankan traversal animasi ketika ada data baru
  useEffect(() => {
  if (!data) return;

  const combined = combineTrees(data); // gabungkan multi tree jadi satu root
  if (!combined) {
    console.warn("No valid trees to combine");
    return;
  }
  const root = { name: combined.name, children: [] };

  setFullTree(combined);
  setTreeData(root);

  treeLiveTraversal(root, combined, setTreeData);
}, [data]);

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

async function treeLiveTraversal(liveRoot, fullNode, setTreeData) {
  const queue = [{ liveNode: liveRoot, fullNode }];

  while (queue.length > 0) {
    const { liveNode, fullNode } = queue.shift();

    if (fullNode.children) {
      for (const child of fullNode.children) {
        const newChild = { name: child.name, children: [] };
        liveNode.children.push(newChild);
        setTreeData(JSON.parse(JSON.stringify(liveRoot)));

        await new Promise((res) => setTimeout(res, 300));
        queue.push({ liveNode: newChild, fullNode: child });
      }
    }
  }
}

function combineTrees(trees) {
  if (!Array.isArray(trees) || trees.length === 0) return null;

  const rootName = trees[0]?.name || "Unknown";
  const recipeGroups = [];

  for (const tree of trees) {
    if (!tree) continue;

    const children = tree.children || [];

    recipeGroups.push({
      name: '',
      children: children
    });
  }

  if (recipeGroups.length === 0) return null;

  return {
    name: rootName,
    children: recipeGroups
  };
}