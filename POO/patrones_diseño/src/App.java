import java.rmi.server.ExportException;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import elements.*;
import elements.operation.*;
import elements.expresion.*;
import visitor.CheckExprVisitor;
import visitor.ExportVisitor;

public class App {
    private static ArrayList<ShapeActions> elements_objt = new ArrayList<ShapeActions>();

    private static void instance_objects(){
        elements_objt.add(new int_exp(17));
        elements_objt.add(new float_exp(9));
        elements_objt.add(new string_exp("Fran"));
    }

    private static void checkXML(){
        CheckExprVisitor exportVisitor = new CheckExprVisitor();

        List<Object> xmlShapes = elements_objt.stream()
        .map(x -> ((ShapeActions<Object>)x).visit(exportVisitor,null))
        .collect(Collectors.toList());

        for (Object s : xmlShapes)
            System.out.println(s);
    }

    private static void xportXML(){
        ExportVisitor exportVisitor = new ExportVisitor();

        List<Object> xmlShapes = elements_objt.stream()
        .map(x -> ((ShapeActions<Object>)x).visit(exportVisitor,null))
        .collect(Collectors.toList());

        for (Object s : xmlShapes)
            System.out.println(s);
    }

    public static void main(String[] args) {
        instance_objects();
        checkXML();
        xportXML();
    }
}

/**
 * ----- RESULTS -----
    Integer
    Float
    Float
    17   
    9.0  
    Fran
 */