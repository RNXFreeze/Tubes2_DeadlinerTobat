'use client';
import RecipeToggle from './RecipeToggle';
import AlgorithmDropdown from './AlgorithmDropdown';
import RecipeSlider from './RecipeSlider';

export default function Sidebar({
  recipeType,
  setRecipeType,
  algorithmType,
  setAlgorithm,
  maxRecipe,
  setMaxRecipe,
  execTime,
  nodeCount,
  totalRecipes,
}) {
  console.log("Sidebar props", { totalRecipes, execTime, nodeCount });
  return (
    <aside className="bg-purple-200 px-4 py-6 space-y-14 min-w-[300px]">
      <div className="flex justify-center text-2xl font-semibold mb-12">
        Recipe Panel
      </div>
      <div><RecipeToggle
        recipeType={recipeType}
        setRecipeType={setRecipeType}
      /></div>
      <div><AlgorithmDropdown
        algorithmType={algorithmType}
        setAlgorithm={setAlgorithm}
      /></div>
      {recipeType === 'multiple' && (
        <div><RecipeSlider value={maxRecipe} setValue={setMaxRecipe} /></div>
      )}
      {execTime !== null && nodeCount !== null &&  totalRecipes !== null && (
        <div className="text-sm text-purple-700 space-y-4">
          <p><strong>Total Recipes:</strong> {totalRecipes} Recipe</p>
          <p><strong>Visited Node:</strong> {nodeCount} Node</p>
          <p><strong>Execution Time:</strong> {execTime} ms</p>
        </div>
      )}
    </aside>
  );
}
