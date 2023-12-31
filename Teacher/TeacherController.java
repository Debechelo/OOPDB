import java.util.List;

public class TeacherController {
    private TeacherService teacherService;
    private TeacherView teacherView;

    public TeacherController(TeacherService teacherService, TeacherView teacherView) {
        this.teacherService = teacherService;
        this.teacherView = teacherView;
    }

    public void createTeacher(int id, String name, String subject) {
        Teacher teacher = new Teacher(id, name, subject);
        teacherService.addTeacher(teacher);
    }

    public void editTeacher(int id, String name, String subject) {
        teacherService.editTeacher(id, name, subject);
    }

    public void displayTeachers() {
        List<Teacher> teachers = teacherService.getTeachers();
        teacherView.printTeachers(teachers);
    }
}