'use client';
import Tree from 'react-d3-tree';
import { useState, useEffect } from 'react';

export default function TreeVisualizer({ target, algorithmType, maxRecipe, setNodeCount, enabled }) {
  const [treeData, setTreeData] = useState(null);

  useEffect(() => {
    if (!enabled || !target || !algorithmType || maxRecipe == null || maxRecipe < 0) return;

    setNodeCount(0);

    const baseURL = 'http://localhost:8080/api';
    const algoPath = algorithmType.toLowerCase();
    const url = `${baseURL}/${algoPath}?target=${encodeURIComponent(target)}&max_recipe=${maxRecipe}`;

    fetch(url)
      .then(res => res.json())
      .then(data => {
        const combined = combineTrees(data.trees, target);
        setTreeData(combined);
        setNodeCount(countNodes(combined));
      })
      .catch(err => console.error('Fetch error:', err));
  }, [target, maxRecipe, algorithmType, setNodeCount]);

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
            if (nodeDatum.name === '') return <g></g>;
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

function combineTrees(trees, rootName = 'Root') {
  if (!Array.isArray(trees) || trees.length === 0) return null;

  const recipePairs = [];

  for (const tree of trees) {
    if (!tree || !Array.isArray(tree.children)) continue;
    recipePairs.push({
      name: '',
      children: tree.children
    });
  }

  if (recipePairs.length === 0) return null;

  return {
    name: rootName,
    children: recipePairs
  };
}

function countNodes(node) {
  if (!node) return 0;
  let count = 1;
  if (node.children) {
    for (const child of node.children) {
      count += countNodes(child);
    }
  }
  return count;
}