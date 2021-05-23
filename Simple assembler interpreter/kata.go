package Simple_assembler_interpreter

import (
	"strings"
	"strconv"
)

func SimpleAssembler(program []string) map[string]int {
	programResult := make(map[string]int)


	for currentPointer := 0; currentPointer < len(program); currentPointer++ {
		statement := program[currentPointer]
		statementPart := strings.Split(statement, " ")
		switch statementPart[0] {
		case "mov":
			v, ok := programResult[statementPart[2]]
			if ok {
				programResult[statementPart[1]] = v
			} else {
				add, _ := strconv.Atoi(statementPart[2])
				programResult[statementPart[1]] = add
			}
		case "inc":
			programResult[statementPart[1]] = programResult[statementPart[1]]+1
		case "dec":
			programResult[statementPart[1]] = programResult[statementPart[1]]-1
		case "jnz":

			v, ok := programResult[statementPart[1]]
			if ok && v != 0 {
				add, _ := strconv.Atoi(statementPart[2])
				currentPointer = currentPointer + add - 1
			} else if !ok && statementPart[1] != "" {
				v, _ := strconv.Atoi(statementPart[2])
				if v != 0 {
					add, _ := strconv.Atoi(statementPart[2])
					currentPointer = currentPointer + add - 1
				}
			}
		}
	}
	return programResult
}
