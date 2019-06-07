package pos

import (
	"bytes"
	"encoding/gob"
)

/* Feature Gob interface */

func (sf singleFeature) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	ee := &errEncoder{enc: encoder.Encode}

	ee.encode(sf.featureType)
	ee.encode(sf.value)
	if ee.err != nil {
		return nil, ee.err
	}
	return buf.Bytes(), nil
}

func (sf *singleFeature) GobDecode(buf []byte) error {
	b := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(b)
	ed := &errDecoder{dec: decoder.Decode}

	ed.decode(&sf.featureType)
	ed.decode(&sf.value)
	return ed.err
}

func (tf tupleFeature) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	ee := &errEncoder{enc: encoder.Encode}

	ee.encode(tf.featureType)
	ee.encode(tf.value1)
	ee.encode(tf.value2)
	if ee.err != nil {
		return nil, ee.err
	}
	return buf.Bytes(), nil
}

func (tf *tupleFeature) GobDecode(buf []byte) error {
	b := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(b)
	ed := &errDecoder{dec: decoder.Decode}

	ed.decode(&tf.featureType)
	ed.decode(&tf.value1)
	ed.decode(&tf.value2)
	return ed.err
}

/* fctuple Gob Interface */
func (fc fctuple) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	ee := &errEncoder{enc: encoder.Encode}

	ee.encode(&fc.feature)
	ee.encode(fc.POSTag)
	if ee.err != nil {
		return nil, ee.err
	}
	return buf.Bytes(), nil
}

func (fc *fctuple) GobDecode(buf []byte) error {
	b := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(b)
	ed := &errDecoder{dec: decoder.Decode}

	ed.decode(&fc.feature)
	ed.decode(&fc.POSTag)
	return ed.err
}

/* Perceptron Gob Interface */

func (p *perceptron) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	ee := &errEncoder{enc: encoder.Encode}
	ee.encode(&p.weightsSF)
	ee.encode(&p.weightsTF)
	ee.encode(&p.totals)
	ee.encode(&p.steps)
	ee.encode(p.instancesSeen)
	if ee.err != nil {
		return nil, ee.err
	}
	return buf.Bytes(), nil
}

type errEncoder struct {
	enc func(e interface{}) error
	err error
}

func (ee *errEncoder) encode(e interface{}) {
	if ee.err != nil {
		return
	}
	ee.err = ee.enc(e)
}

func (p *perceptron) GobDecode(buf []byte) error {
	b := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(b)
	ed := &errDecoder{dec: decoder.Decode}

	ed.decode(&p.weightsSF)
	ed.decode(&p.weightsTF)
	ed.decode(&p.totals)
	ed.decode(&p.steps)
	ed.decode(&p.instancesSeen)
	return ed.err
}

type errDecoder struct {
	dec func(e interface{}) error
	err error
}

func (ed *errDecoder) decode(e interface{}) {
	if ed.err != nil {
		return
	}
	ed.err = ed.dec(e)
}

func init() {
	gob.Register(singleFeature{})
	gob.Register(tupleFeature{})
}
