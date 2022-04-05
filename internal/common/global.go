package common

import "math"

// MinFastSort Minimal Fast Sort
func MinFastSort(x []float64, idx []int, n int, m int) ([]float64, []int) {

	for i := 0; i < m; i++ {
		for j := i + 1; j < n; j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
				idx[i], idx[j] = idx[j], idx[i]
			}
		}
	}
	return x, idx
}

// DistVector calculates the distance of two vector
func DistVector(vectors1 []float64, vectors2 []float64) float64 {

	sum := 0.0
	for index, vec1 := range vectors1 {
		sum += (vec1 - vectors2[index]) * (vec1 - vectors2[index])
	}

	return math.Sqrt(sum)
}

// ObservePointDistVector calculates the distance of two vector
func ObservePointDistVector(vectors1 []float64, vectors2 []float64) float64 {
	sum := 0.0
	for index, vec1 := range vectors1 {
		sum += (vec1 - vectors2[index]) * (vec1 - vectors2[index])
	}

	return math.Sqrt(sum)
}

// Get seed number for random and start it up 
func Randomize()
{
	for i := 0; i <= 54; ++i  {
		this->oldrand[i] = 0.0;
	}

	this->jrand = 0;
	WarmupRandom((unsigned)time(NULL));
	return;
}


// Get randomize off and running 
func WarmupRandom(double seed)
{
	var j1, ii int
	double new_random, prev_random;
	oldrand[54] = seed;
	new_random = 0.000000001;
	prev_random = seed;

	for j1 = 1; j1 <= 54; ++j1	{
		ii = (21*j1)%54;
		oldrand[ii] = new_random;
		new_random = prev_random - new_random;
		if(new_random < 0.0)
			new_random += 1.0;
		prev_random = oldrand[ii];
	}

	AdvanceRandom();
	AdvanceRandom();
	AdvanceRandom();
	jrand = 0;
}

// Create next batch of 55 random numbers
func AdvanceRandom ()
{
	newRandom := 0.0;
	for int j1 = 0 ; j1 < 24 ; j1++ 	{

		newRandom = oldrand[j1]-oldrand[j1+31];
		if( newRandom < 0.0 )	{

			newRandom = newRandom+1.0;
		}
		oldrand[j1] = newRandom;
	}
	for int j1 = 24 ; j1 < 55 ; j1++  {

		newRandom = oldrand[j1]-oldrand[j1-24];
		if( newRandom < 0.0 ){

			newRandom = newRandom+1.0;
		}
		oldrand[j1] = newRandom;
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
