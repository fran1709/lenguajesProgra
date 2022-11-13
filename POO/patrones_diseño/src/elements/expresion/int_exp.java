package elements.expresion;

import elements.ShapeActions;
import visitor.Visitor;

public class int_exp extends Expresion implements ShapeActions{
    int value;

    public int_exp(int pValue){
        super(String.valueOf(pValue));
        this.value = pValue;
    }

    public int getValue() {
        return value;
    }
    public void setValue(int value) {
        this.value = value;
    }

    @Override
    public Object visit(Visitor v, Object ctx) {
        return v.visit_int(this);
    }   
}