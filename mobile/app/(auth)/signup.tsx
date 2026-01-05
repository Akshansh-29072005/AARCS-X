import { useState } from "react";
import { View, Text, TextInput, Pressable, StyleSheet } from "react-native";
import { signup } from "@/lib/api";
import { theme } from "@/constants/theme";

export default function SignupScreen() {
  const [form, setForm] = useState({
    first_name: "",
    last_name: "",
    email: "",
    phone: "",
    semester: 0,
    branch: "",
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  function update(key: string, value: string) {
    setForm(prev => ({ ...prev, [key]: value }));
  }

  async function handleSignup() {
    try {
      setLoading(true);
      setError("");

      // Convert semester to number before sending
      const payload = {
        ...form,
        semester: parseInt(form.semester.toString()) || 0,
      };

      await signup(payload);

    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <View style={styles.container}>
      <Text style={styles.title}>CREATE_ACCOUNT</Text>

      {Object.keys(form).map(key => (
        <TextInput
          key={key}
          placeholder={key.toUpperCase()}
          placeholderTextColor={theme.colors.muted}
          secureTextEntry={key === "password"}
          value={(form as any)[key].toString()}
          onChangeText={v => update(key, v)}
          style={styles.input}
          keyboardType={key === "semester" || key === "phone" ? "numeric" : "default"}
        />
      ))}

      {error ? <Text style={styles.error}>{error}</Text> : null}

      <Pressable style={styles.button} onPress={handleSignup} disabled={loading}>
        <Text style={styles.buttonText}>
          {loading ? "CREATING..." : "EXECUTE_SIGNUP"}
        </Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: theme.colors.bg,
    justifyContent: "center",
    padding: 24,
  },
  system: {
    color: theme.colors.primary,
    textAlign: "center",
    marginBottom: 12,
    letterSpacing: 1,
  },
  title: {
    color: theme.colors.text,
    fontSize: 32,
    fontWeight: "700",
    textAlign: "center",
    marginBottom: 32,
  },
  input: {
    borderWidth: 1,
    borderColor: theme.colors.border,
    borderRadius: 10,
    padding: 14,
    color: theme.colors.text,
    marginBottom: 16,
  },
  button: {
    backgroundColor: theme.colors.primary,
    padding: 16,
    borderRadius: 12,
    alignItems: "center",
    marginTop: 10,
  },
  buttonText: {
    color: "#000",
    fontWeight: "700",
    letterSpacing: 1,
  },
  error: {
    color: "red",
    marginBottom: 10,
    textAlign: "center",
  },
});
