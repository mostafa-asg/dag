package dag

type pipelineResult struct {
	dag *Dag
}

func (result *pipelineResult) Then() *pipelineDSL {
	return &pipelineDSL{
		result.dag,
	}
}

type pipelineDSL struct {
	dag *Dag
}

func (dsl *pipelineDSL) Spawns(tasks ...func()) *spawnsResult {
	dsl.dag.Spawns(tasks...)
	return &spawnsResult{
		dsl.dag,
	}
}

type spawnsResult struct {
	dag *Dag
}

func (result *spawnsResult) Join() *spawnsDSL {
	return &spawnsDSL{
		result.dag,
	}
}

type spawnsDSL struct {
	dag *Dag
}

func (dsl *spawnsDSL) Pipeline(tasks ...func()) *pipelineResult {
	dsl.dag.Pipeline(tasks...)
	return &pipelineResult{
		dsl.dag,
	}
}
