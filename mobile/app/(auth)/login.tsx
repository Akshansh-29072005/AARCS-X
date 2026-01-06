import { useState } from "react";
import { View, Text, TextInput, TouchableOpacity, StyleSheet, ActivityIndicator, Alert, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { useTheme } from "../../context/ThemeContext";
import { UserRole } from "../../lib/api";

type Role = UserRole;

export default function Login() {
  const router = useRouter();
  const { colors } = useTheme();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState<Role>("student");
  const [loading, setLoading] = useState(false);

  const handleLogin = async () => {
    // For now, simple mock login
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      if (role === 'student') {
        router.replace("/(student)");
      } else {
        Alert.alert("Notice", `${role} dashboard is coming soon!`);
      }
    }, 1000);
  };

  const dynamicStyles = getStyles(colors);

  return (
    <View style={dynamicStyles.container}>
      <Text style={dynamicStyles.title}>AARCS-X</Text>
      <Text style={dynamicStyles.subtitle}>Sign in to continue</Text>

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
        <Text style={dynamicStyles.label}>Email Address</Text>
        <TextInput
          style={dynamicStyles.input}
          placeholder="Enter your email"
          placeholderTextColor={colors.muted}
          value={email}
          onChangeText={setEmail}
          autoCapitalize="none"
        />

        <Text style={dynamicStyles.label}>Password</Text>
        <TextInput
          style={dynamicStyles.input}
          placeholder="Enter your password"
          placeholderTextColor={colors.muted}
          value={password}
          onChangeText={setPassword}
          secureTextEntry
        />

        <TouchableOpacity
          style={dynamicStyles.button}
          onPress={handleLogin}
          disabled={loading}
        >
          {loading ? (
            <ActivityIndicator color="#000" />
          ) : (
            <Text style={dynamicStyles.buttonText}>Sign In</Text>
          )}
        </TouchableOpacity>

        <View style={dynamicStyles.footer}>
          <Text style={dynamicStyles.footerText}>Don't have an account? </Text>
          <TouchableOpacity onPress={() => router.push("/(auth)/signup")}>
            <Text style={dynamicStyles.link}>Sign Up</Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
}

const getStyles = (colors: any) => StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.bg,
    padding: 24,
    justifyContent: "center",
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
    color: "#000", // Black text on primary (Cyan) background
  },
  form: {
    gap: 16,
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
  },
  footerText: {
    color: colors.muted,
  },
  link: {
    color: colors.primary,
    fontWeight: "bold",
  },
});
