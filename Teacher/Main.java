public class Main {
    public static void main(String[] args) {
        TeacherService teacherService = new TeacherService();
        TeacherView teacherView = new TeacherView();
        TeacherController teacherController = new TeacherController(teacherService, teacherView);

        teacherController.createTeacher(1, "John Doe", "Math");
        teacherController.createTeacher(2, "Jane Smith", "Physics");

        teacherController.displayTeachers();

        teacherController.editTeacher(1, "John Doe", "Computer Science");

        teacherController.displayTeachers();
    }
}