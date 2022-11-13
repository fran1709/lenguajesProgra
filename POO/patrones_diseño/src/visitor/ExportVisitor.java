package visitor;

import elements.expresion.float_exp;
import elements.expresion.int_exp;
import elements.expresion.string_exp;

public class ExportVisitor implements Visitor{

    @Override
    public Object visit_int(int_exp ctx) {
        // TODO Auto-generated method stub
        return ctx.getValue();
    }

    @Override
    public Object visit_float(float_exp ctx) {
        // TODO Auto-generated method stub
        return ctx.getValue();
    }

    @Override
    public Object visit_string(string_exp ctx) {
        // TODO Auto-generated method stub
        return ctx.getValue();
    }
    
}
