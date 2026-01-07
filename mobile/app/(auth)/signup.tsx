import { useState } from "react";
import { View, Text, TextInput, TouchableOpacity, StyleSheet, ActivityIndicator, ScrollView, Pressable, Alert } from "react-native";
import { useRouter } from "expo-router";
import { useTheme } from "../../context/ThemeContext";
import { UserRole, signup } from "../../lib/api";

type Role = UserRole;

export default function Signup() {
  const router = useRouter();
  const { colors } = useTheme();
  const [loading, setLoading] = useState(false);
  const [role, setRole] = useState<Role>("student");

  // Form State
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [semester, setSemester] = useState("");
  const [branch, setBranch] = useState("");

  const handleSignup = async () => {
    try {
      setLoading(true);

      await signup({
        first_name: firstName,
        last_name: lastName,
        email,
        phone,
        semester: parseInt(semester) || 0,
        branch,
        role,
      });

      Alert.alert("Success", "Account created successfully!", [
        { text: "OK", onPress: () => router.back() }
      ]);
    } catch (error: any) {
      Alert.alert("Error", error.message);
    } finally {
      setLoading(false);
    }
  };

  const dynamicStyles = getStyles(colors);

  return (
    <View style={dynamicStyles.container}>
      <ScrollView contentContainerStyle={dynamicStyles.scrollContent}>
        <Text style={dynamicStyles.title}>Create Account</Text>
        <Text style={dynamicStyles.subtitle}>Join AARCS-X today</Text>

        {/* Role Selection */}
        <View style={dynamicStyles.roleContainer}>
          {(["student", "faculty", "institute"] as Role[]).map((r) => (
            <Pressable
              key={r}
              style={[
                dynamicStyles.roleButton,
                role === r && dynamicStyles.roleButtonActive,
              ]}
              onPress={() => setRole(r)}
            >
              <Text style={[dynamicStyles.roleText, role === r && dynamicStyles.roleTextActive]}>
                {r.charAt(0).toUpperCase() + r.slice(1)}
              </Text>
            </Pressable>
          ))}
        </View>

        <View style={dynamicStyles.form}>
          <View style={dynamicStyles.row}>
            <View style={dynamicStyles.halfInput}>
              <Text style={dynamicStyles.label}>First Name</Text>
              <TextInput
                style={dynamicStyles.input}
                placeholder="John"
                placeholderTextColor={colors.muted}
                value={firstName}
                onChangeText={setFirstName}
              />
            </View>
            <View style={dynamicStyles.halfInput}>
              <Text style={dynamicStyles.label}>Last Name</Text>
              <TextInput
                style={dynamicStyles.input}
                placeholder="Doe"
                placeholderTextColor={colors.muted}
                value={lastName}
                onChangeText={setLastName}
              />
            </View>
          </View>

          <Text style={dynamicStyles.label}>Email Address</Text>
          <TextInput
            style={dynamicStyles.input}
            placeholder="john@example.com"
            placeholderTextColor={colors.muted}
            value={email}
            onChangeText={setEmail}
            autoCapitalize="none"
            keyboardType="email-address"
          />

          <Text style={dynamicStyles.label}>Phone Number</Text>
          <TextInput
            style={dynamicStyles.input}
            placeholder="9876543210"
            placeholderTextColor={colors.muted}
            value={phone}
            onChangeText={setPhone}
            keyboardType="phone-pad"
          />

          {role === "student" && (
            <View style={dynamicStyles.row}>
              <View style={dynamicStyles.halfInput}>
                <Text style={dynamicStyles.label}>Semester</Text>
                <TextInput
                  style={dynamicStyles.input}
                  placeholder="5"
                  placeholderTextColor={colors.muted}
                  value={semester}
                  onChangeText={setSemester}
                  keyboardType="numeric"
                />
              </View>
              <View style={dynamicStyles.halfInput}>
                <Text style={dynamicStyles.label}>Branch</Text>
                <TextInput
                  style={dynamicStyles.input}
                  placeholder="CSE"
                  placeholderTextColor={colors.muted}
                  value={branch}
                  onChangeText={setBranch}
                />
              </View>
            </View>
          )}

          <TouchableOpacity
            style={dynamicStyles.button}
            onPress={handleSignup}
            disabled={loading}
          >
            {loading ? (
              <ActivityIndicator color="#000" />
            ) : (
              <Text style={dynamicStyles.buttonText}>Sign Up</Text>
            )}
          </TouchableOpacity>

          <View style={dynamicStyles.footer}>
            <Text style={dynamicStyles.footerText}>Already have an account? </Text>
            <TouchableOpacity onPress={() => router.back()}>
              <Text style={dynamicStyles.link}>Sign In</Text>
            </TouchableOpacity>
          </View>
        </View>
      </ScrollView>
    </View>
  );
}

const getStyles = (colors: any) => StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.bg,
  },
  scrollContent: {
    padding: 24,
    paddingTop: 60,
  },
  title: {
    fontSize: 32,
    fontWeight: "bold",
    color: colors.primary,
    marginBottom: 8,
    textAlign: "center",
  },
  subtitle: {
    fontSize: 16,
    color: colors.muted,
    marginBottom: 32,
    textAlign: "center",
  },
  roleContainer: {
    flexDirection: "row",
    backgroundColor: colors.card,
    borderRadius: 12,
    marginBottom: 24,
    padding: 4,
  },
  roleButton: {
    flex: 1,
    paddingVertical: 10,
    alignItems: "center",
    borderRadius: 8,
  },
  roleButtonActive: {
    backgroundColor: colors.primary,
  },
  roleText: {
    color: colors.muted,
    fontWeight: "600",
  },
  roleTextActive: {
    color: "#000",
  },
  form: {
    gap: 16,
  },
  row: {
    flexDirection: "row",
    gap: 12,
  },
  halfInput: {
    flex: 1,
    gap: 8, // Gap between label and input
  },
  label: {
    color: colors.text,
    fontSize: 14,
    marginBottom: -8,
  },
  input: {
    backgroundColor: colors.input,
    borderWidth: 1,
    borderColor: colors.border,
    borderRadius: 12,
    padding: 16,
    color: colors.text,
    fontSize: 16,
  },
  button: {
    backgroundColor: colors.primary,
    padding: 16,
    borderRadius: 12,
    alignItems: "center",
    marginTop: 8,
  },
  buttonText: {
    color: "#000",
    fontSize: 16,
    fontWeight: "bold",
  },
  footer: {
    flexDirection: "row",
    justifyContent: "center",
    marginTop: 24,
    marginBottom: 40,
  },
  footerText: {
    color: colors.muted,
  },
  link: {
    color: colors.primary,
    fontWeight: "bold",
  },
});
