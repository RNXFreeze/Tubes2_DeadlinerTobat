export default function RecipeToggle({recipeType, setRecipeType}) {
    return (
        <div className="flex justify-center">
            <div className="inline-flex rounded-lg shadow-sm bg-purple-300" role="group">
                <button
                    onClick={() => setRecipeType('shortest')}
                    className={`px-3 py-2 font-semibold rounded-l-lg ${
                        recipeType === 'shortest' ? 'bg-purple-500' : 'hover:bg-purple-400'
                    }`}
                >
                    Single Recipe
                </button>
                <button
                    onClick={() => setRecipeType('multiple')}
                    className={`px-3 py-2 font-semibold rounded-r-lg ${
                        recipeType === 'multiple' ? 'bg-purple-500' : 'hover:bg-purple-400'
                    }`}>
                    Multiple Recipes
                </button>
            </div>
        </div>
    );
}