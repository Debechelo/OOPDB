package ControlGroup;

import java.util.List;
import University.*;

public class EducationalGroup {
    private Teacher teacher;
    private List<Student> students;

    public EducationalGroup(Teacher teacher, List<Student> students) {
        this.teacher = teacher;
        this.students = students;
    }


    public Teacher getTeacher() {
        return teacher;
    }

    public List<Student> getStudents() {
        return students;
    }
}