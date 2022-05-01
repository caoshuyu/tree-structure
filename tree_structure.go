package treestructure

import (
	"github.com/caoshuyu/tree-structure/protofile/tsproto"
)

// TreeStructure 树结构
type treeStructure struct {
	tsHeader *tsproto.TreeHeader
}

// NewTreeStructure 新建结构体
func NewTreeStructure(header *tsproto.TreeHeader) *treeStructure {
	t := new(treeStructure)
	if header == nil {
		t.tsHeader = new(tsproto.TreeHeader)
	} else {
		t.tsHeader = header
	}
	return t
}

// AddTreeData 添加数据
func (t *treeStructure) AddTreeData(dataList []string) {
	for _, data := range dataList {
		t.tsHeader.Children = t.addData(t.tsHeader.Children, []rune(data))
	}
}

// DelTreeData 删除数据
func (t *treeStructure) DelTreeData(dataList []string) {
	haveKeyList := t.childData(t.tsHeader.Children)
	haveKeyMap := make(map[string]bool)
	for k := range haveKeyList {
		haveKeyMap[haveKeyList[k]] = true
	}
	for _, data := range dataList {
		if _, h := haveKeyMap[data]; !h {
			return
		}
		delete(haveKeyMap, data)
	}
	bodyList := make([]*tsproto.TreeBody, 0)
	for key := range haveKeyMap {
		bodyList = t.addData(bodyList, []rune(key))
	}
	t.tsHeader.Children = bodyList
}

// GetTreeData 获取数据
func (t *treeStructure) GetTreeData() []string {
	return t.childData(t.tsHeader.Children)
}

// GetProto 获取Proto文件
func (t *treeStructure) GetProto() *tsproto.TreeHeader {
	return t.tsHeader
}

func (t *treeStructure) addData(bodyList []*tsproto.TreeBody, dataRune []rune) []*tsproto.TreeBody {
	if len(dataRune) == 0 {
		return bodyList
	}
	isHave := false
	for k, body := range bodyList {
		bodyRune := []rune(body.Body)
		if bodyRune[0] == dataRune[0] {
			isHave = true
			if len(dataRune) > 0 {
				bodyList[k].Children = t.addData(bodyList[k].Children, dataRune[1:])
			}
		}
	}
	if !isHave {
		bodyList = append(bodyList, t.buildDataTree(dataRune))
	}
	return bodyList
}

func (t *treeStructure) buildDataTree(dataRune []rune) *tsproto.TreeBody {
	if len(dataRune) == 0 {
		return nil
	}
	body := new(tsproto.TreeBody)
	body.Body = string(dataRune[0])
	if len(dataRune) > 1 {
		body.Children = append(body.Children, t.buildDataTree(dataRune[1:]))
	}
	return body
}

func (t *treeStructure) childData(children []*tsproto.TreeBody) []string {
	var dataList []string
	for _, one := range children {
		if len(one.Children) > 0 {
			list := t.childData(one.Children)
			for _, val := range list {
				dataList = append(dataList, one.Body+val)
			}
		} else {
			dataList = append(dataList, one.Body)
		}
	}
	return dataList
}
