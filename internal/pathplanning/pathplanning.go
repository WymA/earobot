package pathplanning

import (
	"fmt"
	"time"
)

// total generation numbers
var genNum int = 100

// current generation numbers
var curGenNum int = 0

// population size
var popSize int = 100

// objective numbers
var objNum int = 3

// parent population
var parentPop = Population{}

// old random numbers in 55 batch
var oldRand [55]float64

// random
var jRand int

// path 1 - 3
var path1 []int
var path2 []int
var path3 []int

//
var alreadyStarted bool

//
var alreadyRun bool

// Individual is the subject in Evolution Algorithms
type Individual struct {
	Rank      int       // Ranking mark
	Obj       []float64 // Object list
	CrowdDist float64   // Distance among Objects in a Population
	Path      []int     // Sepecified for PathPlanning
}

// Population is a group of Individuals.
type Population struct {
	Individuals []Individual
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

// Get seed number for random and start it up
func randomize() {

	for i := 0; i <= 54; i++ {
		oldRand[i] = 0.0
	}

	jRand = 0
	warmupRandom(float64(time.Now().Unix()))
	return
}

// Get randomize off and running
func warmupRandom(seed float64) {

	var j, i int
	var new_random, prev_random float64
	oldRand[54] = seed
	new_random = 0.000000001
	prev_random = seed

	for j = 1; j <= 54; j++ {
		i = (21 * j) % 54
		oldRand[i] = new_random
		new_random = prev_random - new_random
		if new_random < 0.0 {
			new_random += 1.0
		}

		prev_random = oldRand[i]
	}

	advanceRandom()
	advanceRandom()
	advanceRandom()
	jRand = 0
}

// Create next batch of 55 random numbers
func advanceRandom() {

	newRandom := 0.0
	for i := 0; i < 24; i++ {

		newRandom = oldRand[i] - oldRand[i+31]
		if newRandom < 0.0 {

			newRandom = newRandom + 1.0
		}
		oldRand[i] = newRandom
	}
	for i := 24; i < 55; i++ {

		newRandom = oldRand[i] - oldRand[i-24]
		if newRandom < 0.0 {

			newRandom = newRandom + 1.0
		}
		oldRand[i] = newRandom
	}
}

// /* Fetch a single random number between 0.0 and 1.0 */
// func randomperc( )
// {
// 	++jrand;
// 	if( jrand >= 55 )	{

// 		jrand = 1;
// 		advance_random();
// 	}
// 	return ( (double)oldrand[jrand] );
// }

// /* Fetch a single random integer between low and high including the bounds */
// func rnd ( int low, int high )
// {
// 	int res;
// 	if (low >= high){

// 		res = low;
// 	}else{

// 		res = low + ( randomperc()*( high - low + 1 ) ) ;
// 		if ( res > high ){

// 			res = high;
// 		}
// 	}
// 	return res;
// }
// /* Fetch a single random real number between low and high including the bounds */
// func rndreal ( double low, double high )
// {
// 	return (low + (high-low)*randomperc());
// }

// init parameter
func initParameter(curPara Parameter) {

	// set population size
	// set cross rate
	// set mutation rate
	// set the number of the total generation

	/// TODO
}

// initPopulation inits the population data structure
func initPopulation(pop *Population) {
	/// TODO
}

// initIndividual inits the individual data structure
func initIndividual(pop *Population) {

}

// removeDuplicatedIndividual is.
func removeDuplicatedIndividual() {

}

// removeDuplicatedGene
func removeDuplicatedGene(pop *Population) {

}

// getBestIndividual
func getBestIndividual() {

}

// generationSelection
func generationSelection(parentPop *Population, childPop *Population) {

}

// tournament
func tournament() {

}

//交叉基因
func GenCross(parent1 *Individual, parent2 *Individual, child1 *Individual, child2 *Individual) {

	// double p = rand()%1000/1000.0;
	// int Parent1Length = parent1->xPath.size(); //第一个以及第二个父个体长度
	// int Parent2Length = parent2->xPath.size(); //

	// child1->xPath.clear();
	// child2->xPath.clear();

	// if ( p < Pc ) {

	// 	//随机产生第一个以及第二个父个体需要交叉的一位,不能是第一位（起点）
	// 	int Parent1CrossPoint, Parent2CrossPoint;

	// 	//这里要注意避免起点
	// 	Parent1CrossPoint = 1+rand()%(Parent1Length-1);
	// 	Parent2CrossPoint = 1+rand()%(Parent2Length-1);

	// 	for (int j = 0; j < Parent1CrossPoint; ++j)
	// 		child1->xPath.push_back( parent1->xPath[j] );
	// 	for (int j = Parent2CrossPoint; j < Parent2Length; ++j)
	// 		child1->xPath.push_back( parent2->xPath[j] );

	// 	for (int j = 0; j < Parent2CrossPoint; ++j)
	// 		child2->xPath.push_back( parent2->xPath[j] );
	// 	for (int j = Parent1CrossPoint; j < Parent1Length; ++j)
	// 		child2->xPath.push_back( parent1->xPath[j] );

	// }else{

	// 	for (int j = 0; j < Parent1Length; ++j)
	// 		child1->xPath.push_back( parent1->xPath[j] );
	// 	for (int j = 0; j < Parent2Length; ++j)
	// 		child2->xPath.push_back( parent2->xPath[j] );
	// }
}

//变异算子
func generationMutation(pop *Population) {
	// for(int i = 0; i < popSize; ++i)
	// 	GenMutationInd(pop->ind[i]);
}

func GenMutationInd(ind *Individual) {
	// double p = rand()%1000/1000.0;
	// if( p < Pm && ind.xPath.size() > 2){

	// 	//随机产生变异的位置，不能为头和尾
	// 	int mutPoint;
	// 	mutPoint = 1 + rand()%(ind.xPath.size() - 2);
	// 	int bSeries = GenIsSeries(ind, mutPoint);
	// 	//若该点与前后两点连续则不变异
	// 	if(bSeries != 2){

	// 		//随机变异出一位数，可以取道起点和终点，这里可能出错
	// 		//变异出的不可以是障碍物，如果是障碍物则不变异
	// 		int newVal = rand()%(outNum+1);
	// 		int x = newVal%chartWidth;
	// 		int y = newVal/chartWidth;

	// 		if(chart[y][x] == 0)
	// 			ind.xPath[mutPoint] = newVal;
	// 	}
	// }
}

//0-不连续，1-与前一个连续，2-与前后都连续
func intGenIsSeries(ind *Individual, idx int) int {
	// int pre, cur, next;
	// pre = ind.xPath[idx-1];
	// cur = ind.xPath[idx];

	// if(idx+1 < ind.xPath.size())
	// 	next = ind.xPath[idx+1];
	// else
	// 	next = ind.xPath[idx];

	// int x0, y0, x1, y1, x2, y2;
	// x0 = pre%chartWidth;
	// y0 = pre/chartWidth;
	// x1 = cur%chartWidth;
	// y1 = cur/chartWidth;
	// x2 = next%chartWidth;
	// y2 = next/chartWidth;

	// int c2p = max(abs(x1-x0), abs(y1-y0));
	// int c2n = max(abs(x2-x1), abs(y2-y1));

	// if((c2n == 1) && (c2p == 1))
	// 	return 2;
	// else if(c2p == 1)
	// 	return 1;
	// else
	// 	return 0;
	return 0
}

//插入算子
func geneInsert(pop *Population) {
	// for(int i = 0; i < popSize; ++i){

	// 	int len = pop->ind[i].xPath.size();
	// 	int insertResult;
	// 	for(int j = 1; j < len; ){

	// 		insertResult = geneInsertInd(pop->ind[i], j);

	// 		if(insertResult == 0)
	// 			j += 1;
	// 		else
	// 			j += 2;
	// 		len = pop->ind[i].xPath.size();

	// 	}
	// }
}

func geneInsertInd(ind *Individual, idx int) int {

	// //如果和前一位连续则不插，插入中值法计算出的点
	// int bSeries = GenIsSeries(ind, idx);
	// if( bSeries == 0 )	{

	// 	int pre = ind.xPath[idx-1];
	// 	int cur = ind.xPath[idx];
	// 	int x0, y0, x1, y1, x2, y2;
	// 	int x3= 0, y3 = 0;//真正要插入的序号
	// 	x1 = pre%chartWidth;
	// 	y1 = pre/chartWidth;
	// 	x2 = cur%chartWidth;
	// 	y2 = cur/chartWidth;
	// 	x0 = (x1 + x2)/2;

	// 	if(abs(y2 - y1) == 1)
	// 		y0 = (y1 + y2 + 1) / 2;
	// 	else
	// 		y0 = (y1 + y2) / 2;

	// 	//非障碍物点，可以插入
	// 	if(chart[y0][x0] == 0){

	// 		x3 = x0;
	// 		y3 = y0;
	// 	}else{//障碍物点，找到离该点最近的非障碍物点插入

	// 		int ySpan = max(y0, (chartHeight - 1 - y0));
	// 		double preMinDist, curMinDist;
	// 		//将最短距离设置为一个大的初值
	// 		preMinDist = curMinDist = MAX_DIST;
	// 		//搜索当前行
	// 		curMinDist = SearchLineNearest(x0, y0, y0, preMinDist, x3);
	// 		//在当前行中找到点

	// 		if(curMinDist < MAX_DIST){

	// 			y3 = y0;
	// 		} else {

	// 			//分别搜索上面的行和下面的行
	// 			for(int j = 0; j <= ySpan; ++j)
	// 			{
	// 				if((y0 - j) >= 0)
	// 				{
	// 					curMinDist = SearchLineNearest(x0, y0, y0 - j, preMinDist, x3);
	// 					if(curMinDist < preMinDist)
	// 					{
	// 						y3 = y0;
	// 						preMinDist = curMinDist;
	// 					}
	// 				}
	// 				if((y0 + j) < chartHeight)
	// 				{
	// 					curMinDist = SearchLineNearest(x0, y0, y0 + j, preMinDist, x3);
	// 					if(curMinDist < preMinDist)
	// 					{
	// 						y3 = y0;
	// 						preMinDist = curMinDist;
	// 					}
	// 				}

	// 				if(curMinDist < MAX_DIST)
	// 					break;

	// 			}
	// 		}
	// 	}

	// 	int ins = x3 + chartWidth*y3;
	// 	//有效的插入点

	// 	if(ins != 0 && ins != ind.xPath[idx]){

	// 		intVec::iterator iter = ind.xPath.begin() + idx;
	// 		ind.xPath.insert(iter, ins);
	// 		return 1;
	// 	}
	// }

	return 0
}

//搜索当前行最短距离
//x0,y0：要替换点的 x,y坐标
//curLin：当前扫描行的行值
//curMin：未扫描该行前最短距离
//minX：该行最短距离点的x坐标
func SearchLineNearest(x0 int, y0 int, curLin int, preMin float64, minX *int) float64 {

	// int xSpan = max( x0, (chartWidth-1-x0) );
	// double tempDist, curMin;//#tempDist stands for temporary distance value?
	// curMin = preMin;

	// for( int k = 0; k <= xSpan; ++k ){

	// 	if( ((x0-k) >= 0) && (chart[curLin][x0 - k] == 0) ){

	// 		tempDist = sqrt((double)(k*k + (curLin-y0)*(curLin-y0)));
	// 		if(curMin > tempDist){

	// 			minX = x0 - k;
	// 			curMin = tempDist;
	// 		}
	// 	}
	// 	if( ((x0+k) < chartWidth) && (chart[curLin][x0+k] == 0) ){

	// 		tempDist = sqrt((double)(k*k + (curLin-y0)*(curLin-y0)));
	// 		if(curMin > tempDist)	{

	// 			minX = x0 + k;
	// 			curMin = tempDist;
	// 		}
	// 	}

	// 	//已找到该行上的最近距离点
	// 	if(curMin < preMin)
	// 		return curMin;

	// }

	// //没找到
	return preMin
}

func getPopBestObj(objectives int, path []int) {

}

//对种群进行评估
func evaluate(pop *Population) {
	// for(int i = 0; i < popSize; ++i)
	// 	EvaluateInd(pop->ind[i], i);
}

//对个体进行评估
func EvaluateInd(ind *Individual, index int) {
	// double A = 0.0, B = 0.0, C = 0.0;
	// int n = ind.xPath.size() ;

	// //#对长度的评价
	// for ( int j = 1; j < n; j++ ){

	//
	// 	int x1,x2,y1,y2;
	// 	x1 = ind.xPath[j-1] % chartWidth ;
	// 	y1 = ind.xPath[j-1] / chartWidth ;
	// 	x2 = ind.xPath[j] % chartWidth ;
	// 	y2 = ind.xPath[j] / chartWidth ;
	// 	A += sqrt( (double)(x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) );//两点之间的距离
	// }

	// //#对个体光滑度的评价
	// for (int j = 1; j < n-1; j++){

	//
	// 	double angle, len1,len2;
	// 	int x1, x2, y1, y2, x3, y3,
	// 		vectorx1, vectory1, vectorx2, vectory2 ;

	// 	x1 = ind.xPath[j-1] % chartWidth;
	// 	y1 = ind.xPath[j-1] / chartWidth;
	// 	x2 = ind.xPath[j] % chartWidth;
	// 	y2 = ind.xPath[j] / chartWidth;
	// 	x3 = ind.xPath[j+1] % chartWidth;
	// 	y3 = ind.xPath[j+1] / chartWidth;

	// 	vectorx1 = x2 - x1 ;
	// 	vectory1 = y2 - y1 ;
	// 	vectorx2 = x3 - x2 ;
	// 	vectory2 = y3 - y2 ;

	// 	len1 = sqrt((double)(vectorx1*vectorx1+vectory1*vectory1));
	// 	len2 = sqrt((double)(vectorx2*vectorx2+vectory2*vectory2));

	// 	angle = (vectorx1*vectorx2+vectory1*vectory2)/(len1*len2);//计算向量的夹角
	// 	angle = acos(angle);
	// 	B += angle;
	// }

	// //#对个体安全性的评价
	// for(int j = 0; j < n-1; ++j){

	// 	//当前点在栅格图上的坐标
	// 	int x1, y1, x2, y2, m = 0;
	// 	double len = 0.0;
	// 	double L = 0.0;
	// 	x1 = ind.xPath[j]%chartWidth;
	// 	y1 = ind.xPath[j]/chartWidth;
	// 	vector<int> vecChart;

	// 	for(int i = 0; i < chartWidth; ++i)
	// 		for(int k = 0; k < chartWidth; ++k)
	// 			if(chart[i][k] == 1)
	// 				vecChart.push_back(i*chartWidth + k);

	// 	for(int i = 0; i < vecChart.size(); ++i){

	// 		x2 = vecChart[i]%chartWidth;
	// 		y2 = vecChart[i]/chartWidth;
	// 		len = sqrt((double)((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)));
	// 		if(check(ind,((y2-1)*chartWidth+x2)) && check(ind,(y2*chartWidth+x2-1)))
	// 			m += 1;
	// 		else if(check(ind,((y2+1)*chartWidth+x2)) && check(ind,(y2*chartWidth+x2-1)))
	// 			m += 1;
	// 		else if(check(ind,((y2+1)*chartWidth+x2)) && check(ind,(y2*chartWidth+x2+1)))
	// 			m += 1;
	// 		else if(check(ind,((y2-1)*chartWidth+x2)) && check(ind,(y2*chartWidth+x2+1)))
	// 			m += 1;
	// 		L += len;
	// 		L = L - m * sqrt((double)(chartWidth*chartWidth));//进行惩罚
	// 		if(L < 0)
	// 			L = -L;

	// 	}
	// 	C = L / vecChart.size();
	// }

	// int D = 0;
	// int curX, curY, nextX, nextY ;
	// double kX, kY, nextk ;
	// int x1, x2, y1, y2 ;
	// int count = 0 ;

	// for ( int i = 0 ; i < n - 1 ; i++ ){

	// 	x1 = ind.xPath[i] % chartWidth;
	// 	y1 = ind.xPath[i] / chartWidth;
	// 	x2 = ind.xPath[i+1] % chartWidth;
	// 	y2 = ind.xPath[i+1] / chartWidth;

	// 	curX = x1 ;
	// 	curY = y1 ;

	// 	do{

	// 		if(x2==x1){

	// 			if(y2>y1) {nextY = curY+1; nextX = curX;}
	// 			else {nextY = curY-1; nextX = curX; }
	// 		}else if(y2==y1){

	// 			if(x2>x1) {nextX = curX+1; nextY = curY;}
	// 			else {nextX = curX-1; nextY = curY; }
	// 		}else{

	// 			if(x2>x1) nextX = curX+1;
	// 			else nextX = curX-1;

	// 			if(y2>y1) nextY = curY+1;
	// 			else nextY = curY-1;

	// 			if(x2<x1) nextX++;
	// 			if(y2<y1) nextY++;

	// 			kX = (nextX-x1-0.5)/(x2-x1);
	// 			kY = (nextY-y1-0.5)/(y2-y1);
	// 			nextk = min(kX,kY);

	// 			nextX = floor(x1+0.5+nextk*(x2-x1));
	// 			nextY = floor(y1+0.5+nextk*(y2-y1));

	// 			if(x2<x1&&kX<=kY) nextX--;
	// 			if(y2<y1&&kY<=kX) nextY--;
	// 		}
	// 		if(chart[nextY][nextX] == 1)//如果通过了障碍物
	// 			D++;
	// 		curX = nextX; curY = nextY;
	// 		count++;

	// 		if(count>1000){

	// 			nextX=x2;
	// 			nextY=y2;
	// 		}
	// 	}while(( !( nextX == x2 && nextY == y2 )));
	// 	count=0;
	// }

	// double len = sqrt((double)( (chartWidth-1) * (chartWidth-1)*2) );
	// fit[index].A = len/A;

	// if(B == 0)
	// 	fit[index].B = 1;
	// else
	// 	fit[index].B = 1/B;

	// if(C == 0)
	// 	fit[index].C = 0;
	// else
	// 	fit[index].C = 1/C;

	// ind.obj[2] = A + 100000*D;//长度
	// ind.obj[1] = B + 100000*D;//平滑度
	// ind.obj[0] = -C + 100000*D;//安全性

	// return;
}

/////////////////////////////////////////////////////////////////////
////#Function to check the ind wether is in the ind or not/////
/////////////////////////////////////////////////////////////////////
func check(ind *Individual, i int) bool {
	// for(int k = 0; k < ind.xPath.size(); ++k){

	// 	if( ind.xPath[k] == i )
	// 		return true;
	// }
	return false
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//// Function to assign rank and crowding distance to a population of size pop_size//
//////////////////////////////////////////////////////////////////////////////////////////////////
func assignRankAndCrowdingDistance(new_pop *Population, popSize int) {
	// int flag,/* i, */end, front_size, rank = 1;
	// list* orig, * cur, * temp1, * temp2;
	// orig = new list;
	// cur = new list;
	// front_size = 0;
	// orig->index = -1;
	// orig->parent = NULL;
	// orig->child = NULL;
	// cur->index = -1;
	// cur->parent = NULL;
	// cur->child = NULL;
	// temp1 = orig;

	// for( int i = 0; i < popSize ; ++i ){

	// 	insert( temp1, i ) ;
	// 	temp1 = temp1->child;
	// }
	// do{

	// 	if( orig->child->child == NULL )	{

	// 		new_pop->ind[orig->child->index].rank = rank;
	// 		new_pop->ind[orig->child->index].crowd_dist = INF;
	// 		break;
	// 	}

	// 	temp1 = orig->child;
	// 	insert(cur, temp1->index);
	// 	front_size = 1;
	// 	temp2 = cur->child;
	// 	temp1 = del(temp1);
	// 	temp1 = temp1->child;

	// 	do{

	// 		temp2 = cur->child;

	// 		do{
	// 			end = 0;
	// 			flag = check_dominance(&(new_pop->ind[temp1->index]), &(new_pop->ind[temp2->index]));
	// 			if(flag == 1)
	// 			{
	// 				insert(orig, temp2->index);
	// 				temp2 = del(temp2);
	// 				--front_size;
	// 				temp2 = temp2->child;
	// 			}
	// 			if (flag == 0)
	// 			{
	// 				temp2 = temp2->child;
	// 			}
	// 			if (flag == -1)
	// 			{
	// 				end = 1;
	// 			}
	// 		}while(end != 1 && temp2 != NULL);

	// 		if (flag == 0 || flag == 1)
	// 		{
	// 			insert (cur, temp1->index);
	// 			front_size++;
	// 			temp1 = del (temp1);
	// 		}
	// 		temp1 = temp1->child;

	// 	}while(temp1 != NULL);

	// 	temp2 = cur->child;
	// 	do{

	// 		new_pop->ind[temp2->index].rank = rank;
	// 		temp2 = temp2->child;
	// 	}while (temp2 != NULL);

	// 	assign_crowding_distance_list (new_pop, cur->child, front_size);
	// 	temp2 = cur->child;

	// 	do{

	// 		temp2 = del (temp2);
	// 		temp2 = temp2->child;
	// 	}while (cur->child !=NULL);

	// 	rank+=1;
	// }

	// while (orig->child!=NULL);

	// free (orig);
	// free (cur);
}

/* Routine to compute crowding distance based on ojbective function values when the population in in the form of a list */
func assign_crowding_distance_list(pop *Population /*list *list,*/, front_size int) {
	// int **obj_array;
	// int *dist;
	// int i,j;
	// list *temp;
	// temp = lst;
	// obj_array = new int*[nobj];
	// dist = new int[front_size];
	// for(i = 0; i < nobj; ++i)
	// 	obj_array[i] = new int[front_size];
	// for(j = 0; j < front_size; ++j)
	// {
	// 	dist[j] = temp->index;
	// 	temp = temp->child;
	// }
	// assign_crowding_distance(pop, dist, obj_array, front_size);
	// delete dist;
	// for(i = 0; i < nobj; ++i)
	// 	delete obj_array[i];
	// delete obj_array;
}

/* Routine to compute crowding distance based on objective function values when the population in in the form of an array */
func assign_crowding_distance_indices(pop *Population, c1 int, c2 int) {
	// int **obj_array;
	// int *dist;
	// int i, j;
	// int front_size;
	// front_size = c2 - c1 +1;
	// obj_array = new int*[nobj];
	// dist = new int[front_size];

	// for(i = 0; i < nobj; ++i)
	// 	obj_array[i] = new int[front_size];

	// for(j = 0; j < front_size; ++j)
	// 	dist[j] = c1++;

	// assign_crowding_distance(pop, dist, obj_array, front_size);
	// delete dist;

	// for(i = 0; i < nobj; ++i)
	// 	delete obj_array[i];

	// delete obj_array;
}

/* Routine to compute crowding distances */
func assign_crowding_distance(pop *Population, dist *int, obj_array *int, front_size int) {
	// int i, j;

	// for(i = 0; i < nobj; ++i)
	// {
	// 	for(j = 0; j < front_size; ++j)
	// 		obj_array[i][j] = dist[j];
	// 	quicksort_front_obj(pop, i, obj_array[i], front_size);
	// }

	// for(j = 0; j < front_size; ++j)
	// 	pop->ind[dist[j]].crowd_dist = 0.0;

	// for(i = 0; i < nobj; ++i)
	// 	pop->ind[obj_array[i][0]].crowd_dist = INF;

	// for(i = 0; i < nobj; ++i)
	// {
	// 	for(j = 1; j < front_size-1; ++j)
	// 	{
	// 		if(pop->ind[obj_array[i][j]].crowd_dist != INF)
	// 		{
	// 			if(pop->ind[obj_array[i][front_size-1]].obj[i] == pop->ind[obj_array[i][0]].obj[i])
	// 				pop->ind[obj_array[i][j]].crowd_dist += 0.0;
	// 			else
	// 				pop->ind[obj_array[i][j]].crowd_dist += (pop->ind[obj_array[i][j+1]].obj[i] - pop->ind[obj_array[i][j-1]].obj[i])/(pop->ind[obj_array[i][front_size-1]].obj[i] - pop->ind[obj_array[i][0]].obj[i]);
	// 		}
	// 	}
	// }

	// for(j = 0; j < front_size; ++j)
	// {
	// 	if(pop->ind[dist[j]].crowd_dist != INF)
	// 		pop->ind[dist[j]].crowd_dist = (pop->ind[dist[j]].crowd_dist)/nobj;
	// }

}

func check_dominance(a *Individual, b *Individual) int {

	// int i, flag1, flag2;
	flag1, flag2 := 0, 0

	// for(i = 0; i < nobj; ++i){

	// 	if(a->obj[i] < b->obj[i] ){

	// 		flag1 = 1;
	// 	}else{

	// 		if( a->obj[i] > b->obj[i] )
	// 			flag2 = 1;
	// 	}
	// }

	// if( flag1 == 1 && flag2 == 0 ){

	// 	return 1;
	// }else{

	if flag1 == 0 && flag2 == 1 {
		return -1
	} else {
		return 0
	}

}

func fill_nondominated_sort(mixedPop *Population, new_pop *Population, popSize int) {
	// int flag, i, j, end, front_size, archieve_size, rank = 1;
	// list *pool, *elite, *temp1, *temp2;
	// pool = new list;
	// elite = new list;
	// front_size = 0;
	// archieve_size = 0;
	// pool->index = -1;
	// pool->parent = NULL;
	// pool->child = NULL;
	// elite->parent = NULL;
	// elite->child = NULL;
	// temp1 = pool;

	// for(i = 0; i < 2*popSize; ++i){

	// 	insert(temp1, i);
	// 	temp1 = temp1->child;
	// }
	// i = 0;

	// do{

	// 	temp1 = pool->child;
	// 	insert(elite, temp1->index);
	// 	front_size = 1;
	// 	temp2 = elite->child;
	// 	temp1 = del(temp1);
	// 	temp1 = temp1->child;

	// 	do{

	// 		temp2 = elite->child;
	// 		if(temp1 == NULL)
	// 			break;
	// 		do{
	// 			end = 0;
	// 			flag = check_dominance(&(mixedPop->ind[temp1->index]), &(mixedPop->ind[temp2->index]));
	// 			if( flag == 1){

	// 				insert(pool, temp2->index);
	// 				temp2 = del(temp2);
	// 				--front_size;
	// 				temp2 = temp2->child;
	// 			}

	// 			if(flag == 0)
	// 				temp2 = temp2->child;
	// 			if(flag == -1)
	// 				end = 1;

	// 		}while (end != 1 && temp2 != NULL) ;

	// 		if(flag == 0 || flag == 1){

	// 			insert(elite, temp1->index);
	// 			++front_size;
	// 			temp1 = del(temp1);
	// 		}
	// 		temp1 = temp1->child;

	// 	}while(temp1 != NULL);

	// 	temp2 = elite->child;
	// 	j = i ;
	// 	if((archieve_size+front_size) <= popSize){

	// 		do{

	// 			copy_ind(&mixedPop->ind[temp2->index], &new_pop->ind[i]);
	// 			new_pop->ind[i].rank = rank;
	// 			++archieve_size;
	// 			temp2 = temp2->child;
	// 			++i;
	// 		}while(temp2 != NULL);

	// 		assign_crowding_distance_indices(new_pop, j, i-1);
	// 		++rank;
	// 	} else {

	// 		crowding_fill(mixedPop, new_pop, i, front_size, elite, popSize);
	// 		archieve_size = popSize;
	// 		for(j = i; j < popSize; ++j)
	// 			new_pop->ind[j].rank = rank;

	// 	}
	// 	temp2 = elite->child;
	// 	do{

	// 		temp2 = del(temp2);
	// 		temp2 = temp2->child;
	// 	}while(elite->child != NULL);

	// }while(archieve_size < popSize);

	// while(pool != NULL){

	// 	temp1 = pool;
	// 	pool = pool->child;
	// 	delete temp1;
	// }

	// while(elite != NULL){

	// 	temp1 = elite;
	// 	elite = elite->child;
	// 	delete temp1;
	// }
	// return;
}

func crowding_fill(mixedPop *Population, new_pop *Population, count int, front_size int /*list *elite,*/, popSize int) {
	// int *dist, i, j ;
	// list *temp;
	// assign_crowding_distance_list(mixedPop, elite->child, front_size);
	// dist = new int[front_size];
	// temp = elite->child;

	// for(j = 0; j < front_size; ++j){

	// 	dist[j] = temp->index;
	// 	temp = temp->child;
	// }

	// quicksort_dist(mixedPop, dist, front_size);
	// for(i = count, j = front_size-1; i < popSize; ++i, --j)
	// 	copy_ind(&mixedPop->ind[dist[j]], &new_pop->ind[i]);

	// delete dist;

	// return;
}

/* Insert an element X into the list at location specified by NODE */
func insert( /*list *node,*/ x int) {
	// list *temp;
	// if(node == NULL){

	// 	exit(1);
	// }

	// temp = new list;
	// temp->index = x;
	// temp->child = node->child;
	// temp->parent = node;

	// if(node->child != NULL){

	// 	node->child->parent = temp;
	// }
	// node->child =temp;
	// return;
}

/// FIXME
// /* Delete the node NODE from the list */
// func del(node *list) list* {
// 	// list *temp;
// 	// if(node == NULL)
// 	// 	exit(1);

// 	// temp = node->parent;
// 	// temp->child = node->child;

// 	// if(temp->child != NULL)
// 	// 	temp->child->parent = temp;

// 	// delete node;
// 	// return temp;
// }

/* Routine to merge two populations into one */
func merge(pop1 *Population, pop2 *Population, pop3 *Population, popSize int) {
	// int i, k;

	// for(i = 0; i < popSize; ++i)
	// 	copy_ind(&(pop1->ind[i]), &(pop3->ind[i]));

	// for(i = 0, k = popSize; i < popSize; ++i, ++k)
	// 	copy_ind(&(pop2->ind[i]), &(pop3->ind[k]));

	// return;
}

func copy_ind(ind1 *Individual, ind2 *Individual) {
	// int i;
	// ind2->rank = ind1->rank;
	// ind2->crowd_dist = ind1->crowd_dist;
	// ind2->xPath.clear();

	// for(i = 0; i < ind1->xPath.size(); ++i)
	// 	ind2->xPath.push_back(ind1->xPath[i]);

	// for(i = 0; i < nobj; ++i)
	// 	ind2->obj[i] = ind1->obj[i];
	// return;
}

func quicksort_front_obj(pop *Population, objcount int, obj_array []int, obj_array_size int) {
	q_sort_front_obj(pop, objcount, obj_array, 0, obj_array_size-1)
}

func q_sort_front_obj(pop *Population, objcount int, obj_array []int, left int, right int) {
	// int index, temp, i, j ;
	// double pivot;

	// if( left < right ){

	// 	index = rnd( left, right ) ;
	// 	temp = obj_array[right] ;
	// 	obj_array[right] = obj_array[index];
	// 	obj_array[index] = temp;
	// 	pivot = pop->ind[obj_array[right]].obj[objcount];
	// 	i = left - 1;
	// 	for(j = left; j <right; ++j){

	// 		if(pop->ind[obj_array[j]].obj[objcount] <= pivot){

	// 			++i ;
	// 			temp = obj_array[j] ;
	// 			obj_array[j] = obj_array[i] ;
	// 			obj_array[i] = temp ;
	// 		}
	// 	}
	// 	index = i + 1;
	// 	temp = obj_array[index];
	// 	obj_array[index] = obj_array[right];
	// 	obj_array[right] = temp;
	// 	q_sort_front_obj(pop, objcount, obj_array, left, index-1);
	// 	q_sort_front_obj (pop, objcount, obj_array, index+1, right);

	// }
}

func quicksort_dist(pop *Population, dist *int, front_size int) {
	//q_sort_dist(pop, dist, 0, front_size-1);
}

func q_sort_dist(pop *Population, dist *int, left int, right int) {
	// 	int index, temp, i, j ;
	// 	double pivot ;
	// 	if ( left < right ){

	// 		index = rnd(left, right);
	// 		temp = dist[right];
	// 		dist[right] = dist[index];
	// 		dist[index] = temp;
	// 		pivot = pop->ind[dist[right]].crowd_dist;
	// 		i = left-1;
	// 		for (j=left; j<right; j++)
	// 		{
	// 			if (pop->ind[dist[j]].crowd_dist <= pivot)
	// 			{
	// 				i+=1;
	// 				temp = dist[j];
	// 				dist[j] = dist[i];
	// 				dist[i] = temp;
	// 			}
	// 		}
	// 		index=i+1;
	// 		temp = dist[index];
	// 		dist[index] = dist[right];
	// 		dist[right] = temp;
	// 		q_sort_dist (pop, dist, left, index-1);
	// 		q_sort_dist (pop, dist, index+1, right);

	// 	}
}

func allocate_pop(pop *Population, size int, objNum int) (indSize int, objSize int) {

	pop.Individuals = make([]Individual, size)
	for i := 0; i < size; i++ {
		pop.Individuals[i].Obj = make([]float64, objNum)
	}

	indSize = len(pop.Individuals)
	if indSize > 0 {
		objSize = len(pop.Individuals[indSize-1].Obj)
	} else {
		objSize = 0
	}

	return
}

func Run() {

	// init
	curGenNum = 0

	childPop := Population{}
	mixedPop := Population{}
	curPara := Parameter{}

	//////////////////////////////////////////////////////////
	//#If population size is too small OR is not times of 4
	//#Then exit
	popSize = curPara.PopSize
	if popSize < 4 || (popSize%4) != 0 {

		fmt.Printf("Population size should be large over 4 and multi-times of 4.")
		//exit(1)
	}

	//#Allocate memory to population with population size
	allocate_pop(&parentPop, popSize, objNum)
	allocate_pop(&childPop, popSize, objNum)
	allocate_pop(&mixedPop, 2*popSize, objNum)
	randomize()

	//#Initialize the parent
	initParameter(curPara)

	initPopulation(&parentPop)
	evaluate(&parentPop)

	//#assign the parent population
	assignRankAndCrowdingDistance(&parentPop, popSize)

	//#find the best Obj path based on 3 objs respectivefully
	getPopBestObj(0, path1)
	getPopBestObj(1, path2)
	getPopBestObj(2, path3)
	alreadyStarted = true
	alreadyRun = true

	//// graphic TODO
	// 	m_myView->Invalidate();
	// 	m_myView->UpdateWindow();

	curGenNum = 1 //#current generation number = 1

	//// Flush to file TODO
	// 	report_pop(&parentPop, fpt1, popSize);

	// 	report_pop(&parentPop, fpt3, popSize);
	// 	fflush(stdout);
	// 	fflush(fpt1);
	// 	fflush(fpt2);
	// 	fflush(fpt3);
	// 	fflush(fpt4);
	// 	////////////////////////////////////////////////////////////////////////////
	// 	//#1st generation done///////////////////////////////////////////////

	for curGenNum = 2; curGenNum <= genNum; curGenNum++ {
		//#Begin with curGenNum=2; A big loop

		////////////////////////////////////////////////////////////////////////////
		//#Evolving//////////////////////////////////////////////////////////////
		generationSelection(&parentPop, &childPop)
		generationMutation(&childPop)
		geneInsert(&childPop)
		removeDuplicatedGene(&childPop)
		evaluate(&childPop)
		merge(&parentPop, &childPop, &mixedPop, popSize)
		fill_nondominated_sort(&mixedPop, &parentPop, popSize)

		////////////////////////////////////////////////////////////////////////////
		//// TODO

		//report_pop(&parentPop, fpt3, popSize)
		//fflush(fpt3)

		getPopBestObj(0, path1)
		getPopBestObj(1, path2)
		getPopBestObj(2, path3)

		//// graphic TODO
		// m_myView->Invalidate();
		// m_myView->UpdateWindow();

		dstTime := time.Now().UnixNano() / time.Hour.Milliseconds()
		srcTime := time.Now().UnixNano() / time.Hour.Milliseconds()
		for (dstTime - srcTime) <= 100 {
			dstTime = time.Now().UnixNano() / time.Hour.Milliseconds()
		}
	}
	curGenNum--

	//// TODO
	// 	report_pop(&parentPop, fpt2, popSize);

	getPopBestObj(0, path1)
	getPopBestObj(1, path2)
	getPopBestObj(2, path3)

	//// graphic TODO
	// 	//======================
	// 	m_myView->Invalidate();
	// 	m_myView->UpdateWindow();

	// 	fflush(stdout);
	// 	fflush(fpt1);
	// 	fflush(fpt2);
	// 	fflush(fpt3);
	// 	fflush(fpt4);
	// 	fclose(fpt1);
	// 	fclose(fpt2);
	// 	fclose(fpt3);
	// 	fclose(fpt4);

}
