import java.util.ArrayList;
import java.util.List;

class TeacherService {
    private List<Teacher> teachers;

    public TeacherService() {
        this.teachers = new ArrayList<>();
    }

    public void addTeacher(Teacher teacher) {
        teachers.add(teacher);
    }

    public List<Teacher> getTeachers() {
        return teachers;
    }

    // Дополнительные методы для редактирования данных учителя
    public void editTeacher(int id, String name, String subject) {
        for (Teacher teacher : teachers) {
            if (teacher.getId() == id) {
                teacher.setName(name);
                teacher.setSubject(subject);
                break;
            }
        }
    }
}