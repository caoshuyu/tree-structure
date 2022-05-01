package treestructure

import (
	"fmt"
	"testing"
)

func TestMakeTreeStructure(t *testing.T) {
	ts := NewTreeStructure(nil)
	ts.AddTreeData([]string{"L_2020050100001"})
	ts.AddTreeData([]string{"L_2020050100002"})
	ts.AddTreeData([]string{"L_2020050100003"})
	ts.AddTreeData([]string{"L_2020050100004"})
	ts.AddTreeData([]string{"L_2020050100005"})
	ts.AddTreeData([]string{"P_2020050100005", "P_2010050100005"})

	dataList := ts.GetTreeData()
	fmt.Println(dataList)
	ts.DelTreeData([]string{"L_2020050100004"})
	dataList = ts.GetTreeData()
	fmt.Println(dataList)
	proto := ts.GetProto()
	fmt.Println(proto.String())
}
