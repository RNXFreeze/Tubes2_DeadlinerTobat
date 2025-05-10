export default function RecipeToggle({recipeType, setRecipeType}) {
    return (
        <div className="inline-flex rounded-lg shadow-sm bg-purple-300" role="group">
            <button
                onClick={() => setRecipeType('shortest')}
                className={`px-4 py-2 font-semibold rounded-l-lg ${
                    recipeType === 'shortest' ? 'bg-purple-400' : 'hover:bg-purple-500'
                }`}
            >
                Shortest Recipe
            </button>
            <button
                onClick={() => setRecipeType('multiple')}
                className={`px-4 py-2 font-semibold rounded-r-lg ${
                    recipeType === 'multiple' ? 'bg-purple-400' : 'hover:bg-purple-500'
                }`}>
                Multiple Recipes
            </button>
        </div>
    );
}