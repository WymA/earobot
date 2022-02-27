package earobot

func main() {
	// //////////////////////////////////////////////////////////
	// 	//#Initialize sth.////////////////////////////////////////
	// 	//////////////////////////////////////////////////////////
	// 	this->curGenNum = 0;

	// 	FILE *fpt1;
	// 	FILE *fpt2;
	// 	FILE *fpt3;
	// 	FILE *fpt4;

	// 	population child_pop;
	// 	population mixed_pop;

	// 	fpt1 = fopen("initial_pop.out", "w");
	// 	fpt2 = fopen("final_pop.out", "w");
	// 	fpt3 = fopen("all_pop.out", "w");
	// 	fpt4 = fopen("params.out","w");

	// 	//////////////////////////////////////////////////////////
	// 	//#If population size is too small OR is not times of 4
	// 	//#Then exit
	// 	this->popSize = m_CurPara.pSize;
	// 	if( this->popSize < 4 || ( this->popSize%4 ) != 0 ){

	// 		exit(1);
	// 	}
	// 	this->nobj = 3;

	// 	////////////////////////////////////////////////////////////////
	// 	//#Do sth. about file output////////////////////////////////

	// 	fprintf(fpt4, "\n Population size = %d", popSize);
	// 	fprintf(fpt4, "\n Number of generations = %d", GenNum);
	// 	fprintf(fpt4, "\n Number of objective functions = %d", nobj);
	// 	fprintf(fpt4, "\n Probability of crossover of real variable = %e", Pc);
	// 	fprintf(fpt4, "\n Probability of mutation of real variable = %e", Pm);
	// 	fprintf(fpt1, "# of objectives = %d\n", nobj);
	// 	fprintf(fpt2, "# of objectives = %d\n", nobj);
	// 	fprintf(fpt3, "# of objectives = %d\n", nobj);
	// 	////////////////////////////////////////////////////////////////

	// 	//#Allocate memory to population with population size
	// 	allocate_memory_pop(&parent_pop, popSize);
	// 	allocate_memory_pop(&child_pop, popSize);
	// 	allocate_memory_pop(&mixed_pop, 2*popSize);
	// 	randomize();

	// 	//#Initialize the parent
	// 	this->Init(m_CurPara);

	// 	InitPop(&parent_pop);
	// 	Evaluate(&parent_pop);

	// 	//#assign the parent population
	// 	assign_rank_and_crowding_distance(&parent_pop, popSize);

	// 	//#find the best Obj path based on 3 objs respectivefully
	// 	GetPopBestObj(0, path1);
	// 	GetPopBestObj(1, path2);
	// 	GetPopBestObj(2, path3);
	// 	m_bAlreadyStarted = true;
	// 	m_bAlreadyRun = true;

	// 	m_myView->Invalidate();
	// 	m_myView->UpdateWindow();

	// 	curGenNum = 1; //#current generation number = 1
	// 	report_pop(&parent_pop, fpt1, popSize);
	// 	fprintf(fpt3, "# gen = 1\n");
	// 	report_pop(&parent_pop, fpt3, popSize);
	// 	fflush(stdout);
	// 	fflush(fpt1);
	// 	fflush(fpt2);
	// 	fflush(fpt3);
	// 	fflush(fpt4);
	// 	////////////////////////////////////////////////////////////////////////////
	// 	//#1st generation done///////////////////////////////////////////////

	// 	int dstTime, srcTime;

	// 	for( curGenNum = 2; curGenNum <= GenNum; ++curGenNum ){
	// 		//#Begin with curGenNum=2; A big loop

	// 		////////////////////////////////////////////////////////////////////////////
	// 		//#Evolving//////////////////////////////////////////////////////////////
	// 		GenSelection(&parent_pop, &child_pop);
	// 		GenMutation(&child_pop);
	// 		GenInsert(&child_pop);
	// 		GenPopDelSame(&child_pop);
	// 		Evaluate(&child_pop);
	// 		merge(&parent_pop, &child_pop, &mixed_pop, popSize);
	// 		fill_nondominated_sort(&mixed_pop, &parent_pop, popSize);

	// 		////////////////////////////////////////////////////////////////////////////

	// 		fprintf(fpt3, "# gen = %d\n", curGenNum);
	// 		report_pop(&parent_pop, fpt3, popSize);
	// 		fflush(fpt3);

	// 		GetPopBestObj(0, path1);
	// 		GetPopBestObj(1, path2);
	// 		GetPopBestObj(2, path3);

	// 		//====================
	// 		m_myView->Invalidate();
	// 		m_myView->UpdateWindow();

	// 		dstTime = srcTime = GetTickCount();
	// 		do{
	// 			dstTime = GetTickCount();
	// 		}while((dstTime-srcTime) <= 100);
	// 	}
	// 	--curGenNum ;

	// 	report_pop(&parent_pop, fpt2, popSize);

	// 	GetPopBestObj(0, path1);
	// 	GetPopBestObj(1, path2);
	// 	GetPopBestObj(2, path3);

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

	// 	deallocate_memory_pop(&child_pop, popSize);
	// 	deallocate_memory_pop(&mixed_pop, 2*popSize);

	// 	return 0;
}
