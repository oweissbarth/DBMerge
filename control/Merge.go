package control

import "github.com/oweissbarth/DBMerge/model"

/*Merge takes two differentials and combines them into a single Differential*/
func Merge(localDiff model.Differential, remoteDiff model.Differential) model.Differential {
	log.Info("Merging two diffs")

	var conflicts []model.Conflict

	//Conflicts between multiple additions
	for i := 0; i < len(localDiff.Additions); i++ {
		for j := 0; j < len(remoteDiff.Additions); j++ {
			l := localDiff.Additions[i]
			r := remoteDiff.Additions[j]
			if l.PrimaryKey == r.PrimaryKey {
				conflicts = append(conflicts, model.Conflict{l.PrimaryKey, l.Content, r.Content, model.ADD, model.ADD, l.Origin, r.Origin})
				localDiff.Additions = append(localDiff.Additions[:i], localDiff.Additions[i+1:]...)
				remoteDiff.Additions = append(remoteDiff.Additions[:j], remoteDiff.Additions[j+1:]...)
				i--
				break
			}
		}
	}

	//Conflicts between deletions and modifications
	for i := 0; i < len(localDiff.Modifications); i++ {
		for j := 0; j < len(remoteDiff.Deletions); j++ {
			l := localDiff.Modifications[i]
			r := remoteDiff.Deletions[j]
			if l.PrimaryKey == r.PrimaryKey {
				conflicts = append(conflicts, model.Conflict{l.PrimaryKey, l.Content, r.Content, model.MOD, model.DEL, l.Origin, r.Origin})
				localDiff.Modifications = append(localDiff.Modifications[:i], localDiff.Modifications[i+1:]...)
				remoteDiff.Deletions = append(remoteDiff.Deletions[:j], remoteDiff.Deletions[j+1:]...)
				i--
				break
			}
		}
	}

	//Conflicts between multiple modifications
	for i := 0; i < len(localDiff.Modifications); i++ {
		for j := 0; j < len(remoteDiff.Modifications); j++ {
			l := localDiff.Modifications[i]
			r := remoteDiff.Modifications[j]
			if l.PrimaryKey == r.PrimaryKey {
				conflicts = append(conflicts, model.Conflict{l.PrimaryKey, l.Content, r.Content, model.MOD, model.MOD, l.Origin, r.Origin})
				localDiff.Modifications = append(localDiff.Modifications[:i], localDiff.Modifications[i+1:]...)
				remoteDiff.Modifications = append(remoteDiff.Modifications[:j], remoteDiff.Modifications[j+1:]...)
				i--
				break
			}
		}
	}

	return model.Differential{
		append(localDiff.Additions, remoteDiff.Additions...),
		append(localDiff.Deletions, remoteDiff.Deletions...),
		append(localDiff.Modifications, remoteDiff.Modifications...),
		conflicts}
}
