package dmaster

var (
  evidence[] row
  values[] data
  mtypes[] mtype
)

func addType(typeid int, name string) {
  mtypes = append(mtypes, newMType(typeid, name))
}