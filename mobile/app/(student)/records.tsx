import { View, Text, StyleSheet, ScrollView } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useTheme } from "../../context/ThemeContext";
import { MOCK_RESULTS } from "../../lib/data";

export default function AcademicRecords() {
    const { colors } = useTheme();
    const dynamicStyles = getStyles(colors);

    return (
        <SafeAreaView style={dynamicStyles.container}>
            <View style={dynamicStyles.header}>
                <Text style={dynamicStyles.title}>Academic Record</Text>
            </View>

            <ScrollView contentContainerStyle={dynamicStyles.scrollContent}>
                {MOCK_RESULTS.map((semesterResult, semIndex) => (
                    <View key={semIndex} style={dynamicStyles.semesterContainer}>
                        <Text style={dynamicStyles.semesterTitle}>Semester {semesterResult.semester}</Text>

                        {semesterResult.exams.map((exam, examIndex) => (
                            <View key={examIndex} style={dynamicStyles.examCard}>
                                <View style={dynamicStyles.examHeader}>
                                    <Text style={dynamicStyles.examName}>{exam.examName}</Text>
                                    {exam.sgpa && <Text style={dynamicStyles.sgpa}>SGPA: {exam.sgpa}</Text>}
                                </View>

                                {exam.subjects.map((subject, subIndex) => (
                                    <View key={subIndex} style={dynamicStyles.subjectRow}>
                                        <Text style={dynamicStyles.subjectName}>{subject.subject}</Text>
                                        <Text style={dynamicStyles.score}>
                                            {subject.score}/{subject.maxScore}
                                            {subject.grade && <Text style={dynamicStyles.grade}> ({subject.grade})</Text>}
                                        </Text>
                                    </View>
                                ))}
                            </View>
                        ))}
                    </View>
                ))}
            </ScrollView>
        </SafeAreaView>
    );
}

const getStyles = (colors: any) => StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.bg,
    },
    header: {
        padding: 16,
        borderBottomWidth: 1,
        borderBottomColor: colors.border,
    },
    title: {
        fontSize: 20,
        fontWeight: "bold",
        color: colors.text,
    },
    scrollContent: {
        padding: 16,
        paddingBottom: 40,
    },
    semesterContainer: {
        marginBottom: 32,
    },
    semesterTitle: {
        fontSize: 22,
        fontWeight: "bold",
        color: colors.primary,
        marginBottom: 16,
        textAlign: "center",
        textTransform: "uppercase",
        letterSpacing: 1,
    },
    examCard: {
        backgroundColor: colors.card,
        borderRadius: 16,
        padding: 16,
        marginBottom: 16,
        borderWidth: 1,
        borderColor: colors.border,
    },
    examHeader: {
        flexDirection: "row",
        justifyContent: "space-between",
        alignItems: "center",
        marginBottom: 12,
        borderBottomWidth: 1,
        borderBottomColor: colors.border,
        paddingBottom: 8,
    },
    examName: {
        fontSize: 18,
        fontWeight: "bold",
        color: colors.secondary,
    },
    sgpa: {
        fontSize: 14,
        fontWeight: "bold",
        color: colors.primary,
        backgroundColor: "rgba(0, 229, 255, 0.1)",
        paddingHorizontal: 8,
        paddingVertical: 4,
        borderRadius: 8,
    },
    subjectRow: {
        flexDirection: "row",
        justifyContent: "space-between",
        paddingVertical: 8,
    },
    subjectName: {
        color: colors.text,
        fontSize: 15,
    },
    score: {
        color: colors.muted,
        fontSize: 15,
        fontWeight: "500",
    },
    grade: {
        color: colors.primary,
        fontWeight: "bold",
    }
});
