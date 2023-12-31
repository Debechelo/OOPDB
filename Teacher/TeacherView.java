import java.util.List;

class TeacherView {
    public void printTeachers(List<Teacher> teachers) {
        for (Teacher teacher : teachers) {
            System.out.println("ID: " + teacher.getId() + ", Name: " + teacher.getName() + ", Subject: " + teacher.getSubject());
        }
    }
}