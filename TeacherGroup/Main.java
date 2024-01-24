import ControlGroup.*;
import University.*;
import java.util.List;

public class Main {
    public static void main(String[] args) {

        Teacher teacher = new Teacher(1, "Bruno Mars");
        List<Student> students = List.of(
            new Student(1, "Student 1"),
            new Student(2, "Student 2"),
            new Student(3, "Student 3")
        );

        EducationalGroupService service = new EducationalGroupService();
        EducationalGroupController controller = new EducationalGroupController(service);
        EducationalGroup group = controller.createEducationalGroup(teacher, students);

        System.out.println("Учебная группа с преподавателем " + group.getTeacher().getName() + " и " +
                group.getStudents().size() + " студентами создана.");
    }
}