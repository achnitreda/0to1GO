package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	graph              map[string][]string
	stack              []string
	paths              = [][]string{}
	pathsGroups        = [][][]string{}
	filterdGroups      = [][][][]string{}
	finalpaths         = [][][]string{}
	finalFilteredPaths [][][]string
	start, end         string
	antCount           int
)

type Room struct {
	name string
}

type Antfarm struct {
	antCount int
	links    map[string][]string
	start    string
	end      string
}

// 0 - Parsing file
func parseInput(filename string) (*Antfarm, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "\r\n", "\n")

	lines := strings.Split(string(contentStr), "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	farm := &Antfarm{
		links: make(map[string][]string),
	}

	//  ant count
	farm.antCount, err = strconv.Atoi(lines[0])
	if err != nil || farm.antCount <= 0 {
		return nil, fmt.Errorf("invalid number of ants")
	}

	//  rooms and links
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		if line == "##start" {
			i++
			if i >= len(lines) {
				return nil, fmt.Errorf("missing start room")
			}
			room, err := parseRoom(lines[i])
			if err != nil {
				return nil, err
			}
			farm.start = room.name
		} else if line == "##end" {
			i++
			if i >= len(lines) {
				return nil, fmt.Errorf("missing end room")
			}
			room, err := parseRoom(lines[i])
			if err != nil {
				return nil, err
			}
			farm.end = room.name
		} else if strings.Contains(line, "-") {
			//  link
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link format")
			}
			if parts[0] == parts[1] {
				fmt.Println("ERROR: invalid data format")
				os.Exit(0)
			}
			farm.links[parts[0]] = append(farm.links[parts[0]], parts[1])
			farm.links[parts[1]] = append(farm.links[parts[1]], parts[0])

		}
	}

	if farm.start == "" || farm.end == "" {
		return nil, fmt.Errorf("missing start or end room")
	}

	return farm, nil
}

func parseRoom(line string) (*Room, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid room format")
	}
	return &Room{name: parts[0]}, nil
}

// - 1 - Search for all valid paths-----------------------------------------
func search(currNode, target string) {
	stack = append(stack, currNode)

	if currNode == target {
		paths = append(paths, append([]string{}, stack...))
		stack = stack[:len(stack)-1]
		return
	}

	for _, neighbor := range graph[currNode] {
		if !isInStack(neighbor) {
			search(neighbor, target)
		}
	}

	stack = stack[:len(stack)-1]
}

func isInStack(node string) bool {
	for _, v := range stack {
		if v == node {
			return true
		}
	}
	return false
}

// - 2 - group paths based on there's no Collision between them-------------
func filterPaths() {
	for i := 0; i < len(paths); i++ {
		group := [][]string{paths[i]}
		for j := 0; j < len(paths); j++ {
			if i != j && noCollision(paths[j], group) {
				group = append(group, paths[j])
			}
		}
		pathsGroups = append(pathsGroups, group)
	}
}

// !!!!! in final len(path)==1, results to paths are tall and intersect with those small len(path) == 1
func noCollision(path []string, group [][]string) bool {
	pathmap := map[string]bool{}
	for _, pathx := range group {
		for _, room := range pathx[1 : len(pathx)-1] {
			pathmap[room] = true
		}
	}

	for _, room := range path[1 : len(path)-1] {
		if pathmap[room] {
			return false
		}
	}
	return true
}

// 3 - filter groups
func filterGroups() {
	// fmt.Println(pathsGroups, "\n###############################")
	for i := 0; i < len(pathsGroups); i++ {
		multigroups := [][][]string{pathsGroups[i]}
		// fmt.Println(multigroups)
		for j := i + 1; j < len(pathsGroups); j++ {
			if len(pathsGroups[j]) == len(pathsGroups[i]) {
				multigroups = append(multigroups, pathsGroups[j])
				pathsGroups = append(pathsGroups[:j], pathsGroups[j+1:]...)
				j--
			}
		}
		filterdGroups = append(filterdGroups, multigroups)
	}
	// fmt.Println(filterdGroups, "\n###############################")
}

// 4 - final paths => final filtered paths :
func finalPaths() {
	for _, group := range filterdGroups {
		//  fmt.Println(group)
		var groupt [][]string
		min := math.MaxInt32
		for _, paths := range group {

			ln := 0
			for _, path := range paths {
				ln += len(path)
			}

			if ln < min {
				groupt = paths
				min = ln
			}
		}
		finalpaths = append(finalpaths, groupt)
	}
}

func Result() {
	finalFilteredPaths = [][][]string{finalpaths[0]}
	max := len(finalpaths[0])
	for i := 1; i < len(finalpaths); i++ {
		if max < len(finalpaths[i]) {
			max = len(finalpaths[i])
			finalFilteredPaths = append(finalFilteredPaths, finalpaths[i])
		}
	}
}

// 5 - final Step :
func chooseBestPath(antCount int, finalFilteredPaths [][][]string) [][]string {
	bestGroup := [][]string{}
	minTurns := math.MaxInt32
	for _, group := range finalFilteredPaths {
		turns := calculateTurns(group, antCount)
		if turns < minTurns {
			minTurns = turns
			bestGroup = group
		}
	}
	// fmt.Println("===> Turns,", minTurns, "===> best group", bestGroup)
	return bestGroup
}

func calculateTurns(group [][]string, antCount int) int {
	pathCosts := make([]int, len(group))
	antsOnPath := make([]int, len(group))

	// kol path ch7al fih dyal rooms o kol ant raha f chi room, ch7al mn ant filkhar ratchit
	// fmt.Println("-----intial start dyal ants ----")
	for i, path := range group {
		pathCosts[i] = len(path) - 2
		antsOnPath[i] = 1
		antCount--
	}
	// fmt.Println("	- pathCosts -> ", pathCosts, "antsOnPath -> ", antsOnPath, "ants ->", antCount)

	// fmt.Println("-----distribution dyal nmal liba9i ----")
	for antCount > 0 {
		minIndex := 0
		for i := 1; i < len(group); i++ {
			// fmt.Println("minIdex :", minIndex, "	- pathCosts -> ", pathCosts[minIndex], "antsOnPath -> ", antsOnPath[minIndex], "sum: ", pathCosts[minIndex]+antsOnPath[minIndex])
			// fmt.Println("index :", i, "	- pathCosts -> ", pathCosts[i], "antsOnPath -> ", antsOnPath[i], "sum: ", pathCosts[i]+antsOnPath[i])
			if pathCosts[i]+antsOnPath[i] < pathCosts[minIndex]+antsOnPath[minIndex] {
				minIndex = i
			}
		}
		// fmt.Println("	- minIndex ->", minIndex)
		// fmt.Println("antsOnpath =>",antsOnPath)
		antsOnPath[minIndex]++
		antCount--
	}

	// fmt.Println("	- antsOnPath -> ", antsOnPath)
	// fmt.Println("---- Calculate the turns ----")
	maxTurns := 0
	for i, path := range group {
		turns := len(path) - 2 + antsOnPath[i]
		// fmt.Println(i, path, "->", turns)
		if turns > maxTurns {
			maxTurns = turns
		}
	}
	// fmt.Println("maxTurns dyal each group ->", maxTurns)
	return maxTurns
}

// utile (sort path)
func sortPaths() {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if len(paths[i]) > len(paths[j]) {
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	filename := os.Args[1]
	farm, err := parseInput(filename)
	if err != nil {
		fmt.Printf("ERROR: invalid data format, %v\n", err)
		return
	}

	start = farm.start
	end = farm.end
	antCount = farm.antCount
	graph = farm.links

	// --- check wach argments are correct before start looping ⛔ ??????
	search(start, end)

	// fmt.Println("------graph--------")
	// for k, v := range graph {
	// 	fmt.Printf("%s -> %v\n", k, v)
	// }

	// fmt.Println("------ sort paths and print it --------")
	sortPaths()
	// for i, v := range paths {
	// 	fmt.Println(i, v)
	// }

	// fmt.Println("------ Group Paths based on no Colision between them --------")
	filterPaths()
	// fmt.Println("-->",pathsGroups)
	// fmt.Println("------ final Groups paths --------")
	filterGroups()
	// fmt.Println(pathsGroups)
	// for i,v :=range filterdGroups {
	// 	fmt.Println(i,"---------------------")
	// 	for j,w :=range v {
	// 		fmt.Println(j,w)
	// 	}
	// }

	// fmt.Println("------ final valid group paths --------")
	finalPaths()
	Result()
	// for i, v := range finalFilteredPaths{
	// 	fmt.Println(i, v)
	// }

	// fmt.Println("------ FINAL STEP --------")
	var bestPath [][]string
	if len(finalFilteredPaths) == 1 {
		// fmt.Println("-> Best path for", antCount, "ants:", finalFilteredPaths[0])
		bestPath = finalFilteredPaths[0]
	} else {
		bestPath = chooseBestPath(antCount, finalFilteredPaths)
		// fmt.Println("Best path for", antCount, "ants:", bestPath)
	}

	// fmt.Println("------ Ant movements --------")
	printAntMovements(bestPath, antCount)
}

// 6 - Ant movements

func printAntMovements(bestPath [][]string, antCount int) {
	// I can use the before but I don't have some of them because of len(g)==1
	// Calculate turns and ant movements
	pathCosts := make([]int, len(bestPath))
	antsOnPath := make([]int, len(bestPath))
	// totalAnts := antCount

	// Distribute ants
	for i, path := range bestPath {
		pathCosts[i] = len(path) - 2
		antsOnPath[i] = 1
		antCount--
	}

	for antCount > 0 {
		minIndex := 0
		for i := 1; i < len(bestPath); i++ {
			if pathCosts[i]+antsOnPath[i] < pathCosts[minIndex]+antsOnPath[minIndex] {
				minIndex = i
			}
		}
		antsOnPath[minIndex]++
		antCount--
	}
	// fmt.Println("======>", antsOnPath, "<============")

	// Print ant movements
	// antPositions := make([]int, totalAnts+1)
	// fmt.Println("antPositions =>", antPositions, len(antPositions))

	// occupiedRooms := make(map[string]int)
	// fmt.Println(antCount)
	// a := 0
	// i := 1
	// res := ""
	// for {

	// 	if a == 9 {break}
	// }
	// fmt.Println(res)
	// fmt.Println("antPositions =>", antPositions, len(antPositions))
	printMoves(bestPath, antsOnPath)
}

func printMoves(paths [][]string, antsPerPath []int) {
	antsOnrooms := make([][]string, len(paths))
	for i := 0; i < len(paths); i++ {
		antsOnrooms[i] = make([]string, len(paths[i]))
	}
	// fmt.Println(len(antsOnrooms[2]))
	n := 1
	for {
		move(antsOnrooms)
		// fmt.Printf("------after move------\n%q\n",antsOnrooms)
		// set ants on rooms
		for i := 0; i < len(antsPerPath); i++ {
			if antsPerPath[i] > 0 {
				antsOnrooms[i][1] = fmt.Sprintf("L%d", n)
				n++
				antsPerPath[i]--
			}
		}
		// fmt.Printf("------after add------\n%q\n",antsOnrooms)
		// print moves
		for i := 0; i < len(antsOnrooms); i++ {
			for j := len(antsOnrooms[i]) - 1; j > 0; j-- {
				if antsOnrooms[i][j] != "" {
					fmt.Print(antsOnrooms[i][j], "-", paths[i][j], " ")
				}
			}
		}

		if roomsAreEmty(antsOnrooms) {
			break
		}
		fmt.Println()
	}
}

func roomsAreEmty(rooms [][]string) bool {
	for i := 0; i < len(rooms); i++ {
		for j := 1; j < len(rooms[i]); j++ {
			if rooms[i][j] != "" {
				return false
			}
		}
	}
	return true
}

func move(antsOnrooms [][]string) {
	for i := 0; i < len(antsOnrooms); i++ {
		for j := len(antsOnrooms[i]) - 1; j > 0; j-- {
			antsOnrooms[i][j] = antsOnrooms[i][j-1]
		}
		antsOnrooms[i][0] = ""
	}
}

/*
-> EX:5 ->  9
0 [[start A0 A1 A2 end] [start B0 B1 E2 D2 D3 end] [start C0 C1 C2 C3 I4 I5 end]] ✅
1 [[start A0 D1 F2 F3 F4 end] [start B0 B1 A1 A2 end] [start G0 G1 G2 G3 G4 D3 end] [start C0 C1 C2 C3 I4 I5 end]]

===> turns 8
L1-A0 L4-B0 L6-C0
L1-A1 L2-A0 L4-B1 L5-B0 L6-C1
L1-A2 L2-A1 L3-A0 L4-E2 L5-B1 L6-C2 L9-B0
L1-end L2-A2 L3-A1 L4-D2 L5-E2 L6-C3 L7-A0 L9-B1
L2-end L3-A2 L4-D3 L5-D2 L6-I4 L7-A1 L8-A0 L9-E2
L3-end L4-end L5-D3 L6-I5 L7-A2 L8-A1 L9-D2
L5-end L6-end L7-end L8-A2 L9-D3
L8-end L9-end

-------------------------------------
-> EX:4 -> 9 || 100 => less than 1.5 minutes || 1000 => < 2.5 minutes
0 [[richard gilfoyle peter] [richard dinish jimYoung peter]] ✅

===> turns 6

L1-gilfoyle L3-dinish
L1-peter L2-gilfoyle L3-jimYoung L5-dinish
L2-peter L3-peter L4-gilfoyle L5-jimYoung L7-dinish
L4-peter L5-peter L6-gilfoyle L7-jimYoung L9-dinish
L6-peter L7-peter L8-gilfoyle L9-jimYoung
L8-peter L9-peter

-------------------------------------
// -> EX:3 -> 4
0 [[0 1 4 5]] ✅

===> turns 6
L1-1
L1-4 L2-1
L1-5 L2-4 L3-1
L2-5 L3-4 L4-1
L3-5 L4-4
L4-5

-------------------------------------
-> EX:2 -> 20
0 [[0 3] [0 1 2 3]] ✅

===> turns 11
L1-3 L2-1
L2-2 L3-3 L4-1
L2-3 L4-2 L5-3 L6-1
L4-3 L6-2 L7-3 L8-1
L6-3 L8-2 L9-3 L10-1
L8-3 L10-2 L11-3 L12-1
L10-3 L12-2 L13-3 L14-1
L12-3 L14-2 L15-3 L16 -1
L14-3 L16-2 L17-3 L18-1
L16-3 L18-2 L19-3
L18-3 L20-3
------------------------------------
-> EX:1 -> 10
0 [[start h n e end] [start t E a m end]]
2 [[start h A c k end] [start t E a m end] [start 0 o n e end]] ✅

===> turns 8
L1-t L2-h L3-0
L1-E L2-A L3-o L4-t L5-h L6-0
L1-a L2-c L3-n L4-E L5-A L6-o L7-t L8-h L9-0
L1-m L2-k L3-e L4-a L5-c L6-n L7-E L8-A L9-o L10-t
L1-end L2-end L3-end L4-m L5-k L6-e L7-a L8-c L9-n L10-E
L4-end L5-end L6-end L7-m L8-k L9-e L10-a
L7-end L8-end L9-end L10-m
L10-end

--------------------------------
// -> EX:0 -> 4
0 [[0 2 3 1]] ✅

===> turns 6
L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
--------------------------------
-> test0 -> 3
0 [[1 3 4 0] [1 2 5 6 0]] ✅

===> turns 4
L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
--------------------------------
-> test2 -> 3
0 [[0 2 1] [0 3 1]] ✅

====> turns 3
L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
------------------------

*/
