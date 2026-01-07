import { useState } from "react";
import { View, Text, StyleSheet, ScrollView, TextInput, TouchableOpacity, Switch, Alert, ActivityIndicator } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useTheme } from "../../context/ThemeContext";
import { MOCK_STUDENT } from "../../lib/data";
import { submitTicket } from "../../lib/api";

export default function Settings() {
    const { colors, isDarkMode, toggleTheme } = useTheme();
    const [description, setDescription] = useState("");
    const [loading, setLoading] = useState(false);

    const handleRaiseTicket = async () => {
        try {
            setLoading(true);
            await submitTicket({
                studentId: MOCK_STUDENT.id,
                issueType: "Attendance Discrepancy",
                description,
            });
            Alert.alert("Success", "Ticket raised successfully! We will look into it.");
            setDescription("");
        } catch (error: any) {
            Alert.alert("Error", error.message);
        } finally {
            setLoading(false);
        }
    };

    const dynamicStyles = getStyles(colors);

    return (
        <SafeAreaView style={dynamicStyles.container}>
            <View style={dynamicStyles.header}>
                <Text style={dynamicStyles.title}>Settings</Text>
            </View>

            <ScrollView contentContainerStyle={dynamicStyles.scrollContent}>

                {/* Profile Card */}
                <View style={dynamicStyles.section}>
                    <Text style={dynamicStyles.sectionHeader}>Profile Information</Text>
                    <View style={dynamicStyles.card}>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Name</Text>
                            <Text style={dynamicStyles.value}>{MOCK_STUDENT.firstName} {MOCK_STUDENT.lastName}</Text>
                        </View>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Email</Text>
                            <Text style={dynamicStyles.value}>{MOCK_STUDENT.email}</Text>
                        </View>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Phone</Text>
                            <Text style={dynamicStyles.value}>{MOCK_STUDENT.phone}</Text>
                        </View>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Semester</Text>
                            <Text style={dynamicStyles.value}>{MOCK_STUDENT.semester}</Text>
                        </View>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Branch</Text>
                            <Text style={dynamicStyles.value}>{MOCK_STUDENT.branch}</Text>
                        </View>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Updated On</Text>
                            <Text style={dynamicStyles.value}>Jan 06, 2026</Text>
                        </View>
                    </View>
                </View>

                {/* Preferences */}
                <View style={dynamicStyles.section}>
                    <Text style={dynamicStyles.sectionHeader}>Preferences</Text>
                    <View style={dynamicStyles.card}>
                        <View style={dynamicStyles.row}>
                            <Text style={dynamicStyles.label}>Dark Mode</Text>
                            <Switch
                                value={isDarkMode}
                                onValueChange={toggleTheme}
                                trackColor={{ false: "#767577", true: colors.primary }}
                                thumbColor={isDarkMode ? "#fff" : "#f4f3f4"}
                            />
                        </View>
                    </View>
                </View>

                {/* Raise Ticket */}
                <View style={dynamicStyles.section}>
                    <Text style={dynamicStyles.sectionHeader}>Support</Text>
                    <View style={dynamicStyles.card}>
                        <Text style={dynamicStyles.cardTitle}>Raise a Ticket</Text>
                        <Text style={dynamicStyles.cardSubtitle}>Found an issue with your attendance? Let us know.</Text>

                        <TextInput
                            style={dynamicStyles.textArea}
                            placeholder="Describe your issue..."
                            placeholderTextColor={colors.muted}
                            multiline
                            numberOfLines={4}
                            value={description}
                            onChangeText={setDescription}
                        />

                        <TouchableOpacity
                            style={dynamicStyles.button}
                            onPress={handleRaiseTicket}
                            disabled={loading}
                        >
                            {loading ? (
                                <ActivityIndicator color="#000" />
                            ) : (
                                <Text style={dynamicStyles.buttonText}>Submit Ticket</Text>
                            )}
                        </TouchableOpacity>
                    </View>
                </View>

                {/* App Info */}
                <View style={dynamicStyles.footer}>
                    <Text style={dynamicStyles.footerText}>App Version 1.0.0</Text>
                    <Text style={dynamicStyles.footerText}>© 2026 AARCS-X</Text>
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
    section: {
        marginBottom: 24,
    },
    sectionHeader: {
        fontSize: 16,
        fontWeight: '600',
        color: colors.muted,
        marginBottom: 12,
        textTransform: "uppercase",
        letterSpacing: 1,
    },
    card: {
        backgroundColor: colors.card,
        borderRadius: 16,
        padding: 16,
        borderWidth: 1,
        borderColor: colors.border,
    },
    row: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        paddingVertical: 12,
        borderBottomWidth: 1,
        borderBottomColor: colors.border,
    },
    label: {
        color: colors.text,
        fontSize: 16,
    },
    value: {
        color: colors.muted,
        fontSize: 16,
    },
    cardTitle: {
        fontSize: 18,
        fontWeight: 'bold',
        color: colors.text,
        marginBottom: 4,
    },
    cardSubtitle: {
        fontSize: 14,
        color: colors.muted,
        marginBottom: 16,
    },
    textArea: {
        backgroundColor: colors.input,
        borderRadius: 12,
        padding: 12,
        color: colors.text,
        height: 100,
        textAlignVertical: 'top',
        borderWidth: 1,
        borderColor: colors.border,
        marginBottom: 16,
    },
    button: {
        backgroundColor: colors.primary,
        padding: 16,
        borderRadius: 12,
        alignItems: 'center',
    },
    buttonText: {
        color: '#000',
        fontWeight: 'bold',
        fontSize: 16,
    },
    footer: {
        alignItems: 'center',
        marginTop: 20,
        marginBottom: 20,
    },
    footerText: {
        color: colors.muted,
        fontSize: 12,
    },
});
