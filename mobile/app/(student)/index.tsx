import { View, Text, StyleSheet, Dimensions, ScrollView } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useTheme } from "../../context/ThemeContext";
import { MOCK_ATTENDANCE } from "../../lib/data";
import { BarChart } from "react-native-chart-kit";

const screenWidth = Dimensions.get("window").width;

export default function StudentHome() {
    const { colors } = useTheme();
    const totalClasses = MOCK_ATTENDANCE.reduce((acc, curr) => acc + curr.total, 0);
    const attendedClasses = MOCK_ATTENDANCE.reduce((acc, curr) => acc + curr.attended, 0);
    const totalPercentage = (attendedClasses / totalClasses) * 100;

    const chartData = {
        labels: MOCK_ATTENDANCE.map((item) => item.subject.substring(0, 3)),
        datasets: [
            {
                data: MOCK_ATTENDANCE.map((item) => item.attended),
            },
        ],
    };

    const dynamicStyles = getStyles(colors);

    return (
        <SafeAreaView style={dynamicStyles.container}>
            <ScrollView contentContainerStyle={dynamicStyles.scrollContent}>
                <View style={dynamicStyles.header}>
                    <Text style={dynamicStyles.greeting}>Welcome, Akshansh!</Text>
                    <Text style={dynamicStyles.subtitle}>Student Dashboard</Text>
                </View>

                {/* Stats Cards */}
                <View style={dynamicStyles.statsContainer}>
                    <View style={dynamicStyles.card}>
                        <Text style={dynamicStyles.cardValue}>{totalClasses}</Text>
                        <Text style={dynamicStyles.cardLabel}>Total Classes</Text>
                    </View>
                    <View style={[dynamicStyles.card, { borderColor: colors.primary }]}>
                        <Text style={[dynamicStyles.cardValue, { color: colors.primary }]}>
                            {totalPercentage.toFixed(1)}%
                        </Text>
                        <Text style={dynamicStyles.cardLabel}>Attendance</Text>
                    </View>
                </View>

                {/* Chart Section */}
                <View style={dynamicStyles.chartContainer}>
                    <Text style={dynamicStyles.sectionTitle}>Attendance by Subject</Text>
                    <BarChart
                        data={chartData}
                        width={screenWidth - 32}
                        height={220}
                        yAxisLabel=""
                        yAxisSuffix=""
                        chartConfig={{
                            backgroundColor: colors.card,
                            backgroundGradientFrom: colors.card,
                            backgroundGradientTo: colors.card,
                            decimalPlaces: 0,
                            color: (opacity = 1) => `rgba(${parseInt(colors.primary.slice(1, 3), 16)}, ${parseInt(colors.primary.slice(3, 5), 16)}, ${parseInt(colors.primary.slice(5, 7), 16)}, ${opacity})`,
                            labelColor: (opacity = 1) => colors.text,
                            barPercentage: 0.7,
                        }}
                        style={dynamicStyles.chart}
                        showValuesOnTopOfBars
                    />
                </View>

                {/* Detailed List */}
                <View style={dynamicStyles.detailsContainer}>
                    <Text style={dynamicStyles.sectionTitle}>Subject Breakdown</Text>
                    {MOCK_ATTENDANCE.map((item, index) => (
                        <View key={index} style={dynamicStyles.row}>
                            <Text style={dynamicStyles.subjectName}>{item.subject}</Text>
                            <Text style={dynamicStyles.subjectStat}>{item.attended}/{item.total} ({item.percentage}%)</Text>
                        </View>
                    ))}
                </View>

            </ScrollView>
        </SafeAreaView>
    );
}

const getStyles = (colors: any) => StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: colors.bg,
    },
    scrollContent: {
        padding: 16,
    },
    header: {
        marginBottom: 24,
    },
    greeting: {
        fontSize: 24,
        fontWeight: "bold",
        color: colors.text,
    },
    subtitle: {
        fontSize: 16,
        color: colors.muted,
        marginTop: 4,
    },
    statsContainer: {
        flexDirection: "row",
        justifyContent: "space-between",
        marginBottom: 24,
    },
    card: {
        width: "48%",
        backgroundColor: colors.card,
        padding: 20,
        borderRadius: 16,
        borderWidth: 1,
        borderColor: colors.border,
        alignItems: "center",
    },
    cardValue: {
        fontSize: 28,
        fontWeight: "bold",
        color: colors.text,
    },
    cardLabel: {
        fontSize: 14,
        color: colors.muted,
        marginTop: 4,
    },
    chartContainer: {
        marginBottom: 24,
    },
    sectionTitle: {
        fontSize: 18,
        fontWeight: "bold",
        color: colors.text,
        marginBottom: 16,
    },
    chart: {
        borderRadius: 16,
    },
    detailsContainer: {
        backgroundColor: colors.card,
        borderRadius: 16,
        padding: 16,
    },
    row: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        paddingVertical: 12,
        borderBottomWidth: 1,
        borderBottomColor: colors.border,
    },
    subjectName: {
        color: colors.text,
        fontSize: 14,
    },
    subjectStat: {
        color: colors.muted,
        fontSize: 14,
    }
});
