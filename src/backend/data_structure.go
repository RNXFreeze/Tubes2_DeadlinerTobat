/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/data_structure.go              */
/* Deskripsi  : F07 - Data Structure (Node & Tree & Mapping)                     */
/* PIC F07    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
)

type Element struct {
	Name    string;
	Tier    int;
	Parents [][]string;
}

type Gallery struct {
	GalleryName map[string]*Element;
}

type RecipeNode struct {
	Name    string        `json:"name"`;
	Parents []*RecipeNode `json:"children,omitempty"`;
}

type PartialTree struct {
    tree *RecipeNode;
    leaf []*RecipeNode;
}

type AlgorithmResult struct {
	Trees        []*RecipeNode;
	VisitedCount int;
}

var base_element = map[string]struct{} {
	"Water" : {},
	"Earth" : {},
	"Fire"  : {},
	"Time"  : {},
	"Air"   : {},
}

func IsBase(name string) bool {
	_ , check := base_element[name];
	return check;
}

func FindPointer(root *RecipeNode , name string) *RecipeNode {
    if (root == nil) {
        return nil;
    } else if (root.Name == name && len(root.Parents) == 0) {
        return root;
    } else {
		for _ , parent := range root.Parents {
			if node := FindPointer(parent , name) ; node != nil {
				return node;
			}
		}
		return nil;
	}
}

func IsExpandable(element *Element) bool {
    return element != nil && element.Tier > 0 && len(element.Parents) > 0;
}

func GetMidTier(gallery *Gallery , target string) int {
    return GetTier(gallery , target) / 2;
}

func CloneNode(node *RecipeNode) *RecipeNode {
	if (node == nil) {
		return nil;
	} else {
		parents := make([]*RecipeNode , len(node.Parents));
		for idx , cp := range node.Parents {
			parents[idx] = CloneNode(cp);
		}
		return &RecipeNode{Name : node.Name , Parents : parents};
	}
}

func CloneSlice(src []*RecipeNode) []*RecipeNode {
	cp := make([]*RecipeNode , len(src));
	for idx , node := range src {
		cp[idx] = CloneNode(node);
	}
	return cp;
}

func SignatureTree(node *RecipeNode) string {
    if (node == nil) {
        return "";
    } else if len(node.Parents) == 0 {
        return node.Name;
    } else {
		l := SignatureTree(node.Parents[0]);
		r := SignatureTree(node.Parents[1]);
		return node.Name + "(" + l + "," + r + ")";
	}
}

func CloneTreeMap(original *RecipeNode) (*RecipeNode , map[*RecipeNode]*RecipeNode) {
    clone_map := make(map[*RecipeNode]*RecipeNode);
    var CloneRecursive func(*RecipeNode) *RecipeNode;
    CloneRecursive = func(node *RecipeNode) *RecipeNode {
        if (node == nil) {
            return nil;
        } else {
			clone := &RecipeNode{Name : node.Name};
			clone_map[node] = clone;
			if (len(node.Parents) == 2) {
				clone.Parents = []*RecipeNode{CloneRecursive(node.Parents[0]) , CloneRecursive(node.Parents[1])};
			}
			return clone;
		}
    }
    return CloneRecursive(original) , clone_map;
}

func (num *RecipeNode) Marshal() ([]byte , error) {
	return json.MarshalIndent(num , "" , "  ");
}

func Transform(array []any) [][]string {
	var res [][]string;
	for _ , pairs := range array {
		var pair []string;
		for _ , s := range pairs.([]any) {
			pair = append(pair , s.(string));
		}
		res = append(res , pair);
	}
	return res;
}

func GetTier(gallery *Gallery , name string) int {
    if element , check := gallery.GalleryName[name] ; check {
        return element.Tier;
    } else {
		return 0;
	}
}

func CalculateTier(name string , gallery *Gallery , visited map[string]bool) int {
	if (IsBase(name)) {
		return 0;
	} else {
		element := gallery.GalleryName[name];
		if (element == nil) {
			return 1;
		} else if (element.Tier >= 0) {
			return element.Tier;
		} else if (visited[name]) {
			return 1;
		}
		best := 0
		visited[name] = true;
		for _ , p := range element.Parents {
			t1 := CalculateTier(p[0] , gallery , visited);
			t2 := CalculateTier(p[1] , gallery , visited);
			if t := max(t1 , t2) + 1 ; t > best {
				best = t;
			}
		}
		visited[name] = false;
		element.Tier = best;
		return best;
	}
}

func (gallery *Gallery) GetAllNames() []string {
	names := make([]string , 0 , len(gallery.GalleryName));
	for name := range gallery.GalleryName {
		names = append(names , name);
	}
	return names;
}

func LoadRecipeGallery(path string) (*Gallery , error) {
    file , err := os.Open(path);
    if (err != nil) {
        return nil , err;
    } else {
        defer file.Close();
        raw := map[string]any{};
        if err := json.NewDecoder(file).Decode(&raw) ; err != nil {
            return nil , fmt.Errorf("decode JSON : %w" , err);
        } else {
            gallery := &Gallery{GalleryName : map[string]*Element{}};
            check := false;
            for key := range raw {
                if (strings.HasPrefix(strings.ToLower(key), "tier")) {
                    check = true;
                    break;
                }
            }
            if (check) {
                for key , value := range raw {
                    if (strings.HasPrefix(strings.ToLower(key) , "tier")) {
                        str := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(key) , "tier"));
                        num , _ := strconv.Atoi(str);
                        inner , check := value.(map[string]any);
                        if (check) {
                            for name , arr := range inner {
                                var parents [][]string;
                                if (arr != nil) {
                                    parents = Transform(arr.([]any));
                                }
                                gallery.GalleryName[name] = &Element{
                                    Name    : name,
                                    Tier    : num,
                                    Parents : parents,
                                };
                            }
                        }
                    }
                }
            } else {
                for name , arr := range raw {
                    var parents [][]string;
                    if arr != nil {
                        parents = Transform(arr.([]any));
                    }
                    gallery.GalleryName[name] = &Element{
                        Name    : name,
                        Tier    : -1,
                        Parents : parents,
                    };
                }
            }
            for base := range base_element {
                if _ , check := gallery.GalleryName[base] ; !check {
                    gallery.GalleryName[base] = &Element{Name : base , Tier : 0 , Parents : nil};
                } else {
                    gallery.GalleryName[base].Tier = 0;
                }
            }
            for name , element := range gallery.GalleryName {
                if (element.Tier < 0) {
                    CalculateTier(name , gallery , map[string]bool{});
                }
            }
            return gallery , nil;
        }
    }
}