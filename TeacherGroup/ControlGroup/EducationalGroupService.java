package ControlGroup;

import java.util.List;
import University.*;

public class EducationalGroupService {
    public EducationalGroup createEducationalGroup(Teacher teacher, List<Student> students) {
        return new EducationalGroup(teacher, students);
    }
}
