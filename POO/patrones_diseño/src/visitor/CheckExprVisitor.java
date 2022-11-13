package visitor;

import elements.expresion.*;

public class CheckExprVisitor implements Visitor{

    @Override
    public Object visit_int(int_exp ctx) {
        return "Integer";
    }

    @Override
    public Object visit_float(float_exp ctx) {
        // TODO Auto-generated method stub
        return "Float";
    }

    @Override
    public Object visit_string(string_exp ctx) {
        // TODO Auto-generated method stub
        return "Float";
    }
    
}
