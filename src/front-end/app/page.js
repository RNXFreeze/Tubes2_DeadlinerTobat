'use client'

import { useState } from 'react';
import Navbar from '../components/Navbar';
import RecipeToggle from '../components/RecipeToggle';
import AlgorithmDropdown from '../components/AlgorithmDropdown';
import RecipeSlider from '../components/RecipeSlider';


export default function Home() {
  const [recipeType, setRecipeType] = useState('multiple');
  const [algorithmType, setAlgorithm] = useState('BFS');
  const [element,setElement] = useState(' ');
  const [maxRecipe, setMaxRecipe] = useState('10');

  return (
    <div className="bg-purple-100 min-h-screen">
      <Navbar />
      <main className="flex flex-col items-center justify-center p-6 gap-6 text-center">
        <h1 
          className="text-3xl font-bold text-center">
            Little Alchemy Recipe Finder
        </h1>
        <RecipeToggle recipeType={recipeType} setRecipeType={setRecipeType} />
        <AlgorithmDropdown
          algorithmType={algorithmType}
          setAlgorithm={setAlgorithm}
          element={element}
          setElement={setElement}
          handleSearch={() => {
            console.log("Search clicked:", element, algorithmType);
          }}
        />
      {recipeType === 'multiple' && (
        <RecipeSlider value={maxRecipe} setValue={setMaxRecipe} />
      )}
      </main>
    </div>
  );
}