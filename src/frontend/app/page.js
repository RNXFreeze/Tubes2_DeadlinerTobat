'use client';

import { useEffect, useState } from 'react';
import Navbar from '../components/Navbar';
import Sidebar from '../components/Sidebar';
import TreePage from '../components/TreePage';

export default function Home() {
  const [recipeType, setRecipeType] = useState('shortest');
  const [algorithmType, setAlgorithm] = useState('BFS');
  const [maxRecipe, setMaxRecipe] = useState(10);
  const [searchElement, setSearchElement] = useState('');
  const [execTime, setExecTime] = useState(null);
  const [nodeCount, setNodeCount] = useState(null);
  const [treeData, setTreeData] = useState([]); 

  // Dummy values
  // const execTime = 1;
  // const nodeCount = 1;

  return (
    <div className="min-h-screen bg-purple-100 flex flex-col">
      <Navbar />
      <div className="flex flex-grow">
        <Sidebar
          recipeType={recipeType}
          setRecipeType={setRecipeType}
          algorithmType={algorithmType}
          setAlgorithm={setAlgorithm}
          maxRecipe={maxRecipe}
          setMaxRecipe={setMaxRecipe}
          execTime={execTime}
          nodeCount={nodeCount}
        />
        <TreePage 
          algorithmType={algorithmType}
          searchElement={searchElement}
          setSearchElement={setSearchElement}
          execTime={execTime}
          setExecTime={setExecTime}
          nodeCount={nodeCount}
          setNodeCount={setNodeCount}
          maxRecipe={maxRecipe}
          treeData={treeData}
          setTreeData={setTreeData}
        />
      </div>
    </div>
  );
}
