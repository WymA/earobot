package common

// Individual is the subject in Evolution Algorithms
type Individual struct {
	Rank      int       // Ranking mark
	Obj       []float64 // Object list
	CrowdDist float64   // Distance among Objects in a Population
	Path      []int     // Sepecified for PathPlanning
}

// Population is a group of Individuals.
type Population struct {
	ind []Individual
}

type Fitness struct {
	A float64 //对个体长度的评价
	B float64 //对个体光滑度的评价
	C float64 //对个体安全性的评价
	//	 D float64;//综合评价
}

//static struct  Fitness fit[1000];
//static struct Fitness F;

// type  Lists struct {
// 	int index;
// 	lists *parent;
// 	lists *child;
// }

type Parameter struct {
	PopSize int
	T       int
	PropC   float64
	PropM   float64
	Width   int
	Height  int
	Length  bool
	Smooth  bool
	Safe    bool
}
