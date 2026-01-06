
export interface StudentProfile {
    id: string;
    firstName: string;
    lastName: string;
    email: string;
    phone: string;
    semester: number;
    branch: string;
    avatarUrl?: string;
}

export interface AttendanceStat {
    subject: string;
    attended: number;
    total: number;
    percentage: number;
}

export interface SemesterResult {
    semester: number;
    exams: {
        examName: string; // "Class Test-1", "Class Test-2", "End Semester"
        subjects: {
            subject: string;
            score: number;
            maxScore: number;
            grade?: string;
        }[];
        sgpa?: number;
    }[];
}

export const MOCK_STUDENT: StudentProfile = {
    id: "STU123456",
    firstName: "Akshansh",
    lastName: "Sharma",
    email: "akshansh@example.com",
    phone: "9876543210",
    semester: 5,
    branch: "Computer Science Engineering",
};

export const MOCK_ATTENDANCE: AttendanceStat[] = [
    { subject: "Data Structures", attended: 28, total: 32, percentage: 87.5 },
    { subject: "Algorithms", attended: 24, total: 30, percentage: 80.0 },
    { subject: "Database Systems", attended: 20, total: 25, percentage: 80.0 },
    { subject: "OS", attended: 15, total: 20, percentage: 75.0 },
    { subject: "Computer Networks", attended: 18, total: 20, percentage: 90.0 },
];

export const MOCK_RESULTS: SemesterResult[] = [
    {
        semester: 1,
        exams: [
            {
                examName: "Class Test-1",
                subjects: [
                    { subject: "Maths-I", score: 25, maxScore: 30 },
                    { subject: "Physics", score: 22, maxScore: 30 },
                ],
            },
            {
                examName: "Class Test-2",
                subjects: [
                    { subject: "Maths-I", score: 28, maxScore: 30 },
                    { subject: "Physics", score: 24, maxScore: 30 },
                ],
            },
            {
                examName: "End Semester",
                subjects: [
                    { subject: "Maths-I", score: 85, maxScore: 100, grade: "A" },
                    { subject: "Physics", score: 78, maxScore: 100, grade: "B+" },
                ],
                sgpa: 8.5,
            },
        ],
    },
    {
        semester: 2,
        exams: [
            {
                examName: "Class Test-1",
                subjects: [
                    { subject: "Maths-II", score: 20, maxScore: 30 },
                    { subject: "Chemistry", score: 26, maxScore: 30 },
                ],
            },
            {
                examName: "End Semester",
                subjects: [
                    { subject: "Maths-II", score: 80, maxScore: 100, grade: "A" },
                    { subject: "Chemistry", score: 82, maxScore: 100, grade: "A" },
                ],
                sgpa: 8.2,
            },
        ],
    },
];
