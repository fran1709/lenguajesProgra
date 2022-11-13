package elements;

import visitor.Visitor;

public interface ShapeActions<T> {
    public T visit(Visitor v, T ctx);
}