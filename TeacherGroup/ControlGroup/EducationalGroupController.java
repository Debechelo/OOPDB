package ControlGroup;
import java.util.List;

import University.*;

public class EducationalGroupController {
    private EducationalGroupService groupService;

    public EducationalGroupController(EducationalGroupService groupService) {
        this.groupService = groupService;
    }

    public EducationalGroup createEducationalGroup(Teacher teacher, List<Student> students) {
        return groupService.createEducationalGroup(teacher, students);
    }
}