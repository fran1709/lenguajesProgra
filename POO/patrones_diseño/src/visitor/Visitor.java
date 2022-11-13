package visitor;

import elements.expresion.*;

public interface Visitor<T> {
    public T visit_int(int_exp ctx);
    public T visit_float(float_exp ctx);
    public T visit_string(string_exp ctx);
}
