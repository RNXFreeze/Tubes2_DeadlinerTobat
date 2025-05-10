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
}) {
  return (
    <aside className="bg-purple-200 px-4 py-6 space-y-16 min-w-[300px]">
      <div className="flex justify-center text-2xl font-semibold mb-12">
        Search Panel
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
      {execTime !== null && nodeCount !== null && (
        <div className="text-sm text-purple-700 space-y-4">
          <p><strong>Execution time:</strong> {execTime}ms</p>
          <p><strong>Nodes visited:</strong> {nodeCount}</p>
        </div>
      )}
    </aside>
  );
}
