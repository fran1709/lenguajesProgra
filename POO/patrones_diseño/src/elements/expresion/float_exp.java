package elements.expresion;

import elements.ShapeActions;
import visitor.Visitor;

public class float_exp extends Expresion implements ShapeActions{
    float value;

    public float_exp(float pValue) {
        super(String.valueOf(pValue));
        this.value = pValue;
    }

    public float getValue() {
        return value;
    }
    public void setValue(float value) {
        this.value = value;
    }

    @Override
    public Object visit(Visitor v, Object ctx) {
        return v.visit_float(this);
    }
}