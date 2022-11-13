package elements.expresion;

import elements.ShapeActions;
import visitor.Visitor;

public class string_exp extends Expresion implements ShapeActions{
    String value;

    public string_exp(String pValue){
        super(pValue);
        this.value = pValue;
    }

    public String getValue() {
        return value;
    }
    public void setValue(String value) {
        this.value = value;
    }

    @Override
    public Object visit(Visitor v, Object ctx) {
        return v.visit_string(this);
    }
}
