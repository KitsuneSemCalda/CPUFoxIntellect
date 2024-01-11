package mathutils

type DataPoint struct {
	Class      int
	Attributes []float32
}

func NewDataPoint(class int, attributes []float32) *DataPoint {
	if class < 0 {
		class *= -1
	}
	for i, attr := range attributes {
		if attr < 0 {
			attributes[i] = attr * -1
		}
	}
	return &DataPoint{
		Class:      class,
		Attributes: attributes,
	}
}

func (d *DataPoint) AddAttribute(attr float32) {
	if attr < 0 {
		attr *= -1
	}
	d.Attributes = append(d.Attributes, attr)
}

func (d *DataPoint) RemoveAttribute(index int) {
	if index >= 0 && index < len(d.Attributes) {
		d.Attributes = append(d.Attributes[:index], d.Attributes[index+1:]...)
	}
}

func (d *DataPoint) UpdateClass(class int) {
	if class < 0 {
		class *= -1
	}
	d.Class = class
}

func (d *DataPoint) GetClass() int {
	return d.Class
}

func (d *DataPoint) GetAttributes() []float32 {
	return d.Attributes
}
