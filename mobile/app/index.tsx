import { View, Text, Pressable, StyleSheet } from "react-native";
import { useRouter } from "expo-router";
import { theme } from "@/constants/theme";

export default function Index() {
  const router = useRouter();

  return (
    <View style={styles.container}>
      <Text style={styles.system}>SYSTEM.INITIALIZE(USER)</Text>

      <Text style={styles.title}>AARCS-X</Text>
      <Text style={styles.subtitle}>AUTH_GATEWAY</Text>

      <Pressable
        style={styles.primaryButton}
        onPress={() => router.push("/login")}
      >
        <Text style={styles.primaryText}>EXECUTE_LOGIN</Text>
      </Pressable>

      <Pressable
        style={styles.secondaryButton}
        onPress={() => router.push("/signup")}
      >
        <Text style={styles.secondaryText}>CREATE_ACCOUNT</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: theme.colors.bg,
    justifyContent: "center",
    alignItems: "center",
    padding: 24,
  },
  system: {
    color: theme.colors.primary,
    letterSpacing: 1,
    marginBottom: 8,
  },
  title: {
    color: theme.colors.text,
    fontSize: 36,
    fontWeight: "800",
    letterSpacing: 2,
  },
  subtitle: {
    color: theme.colors.muted,
    marginBottom: 40,
    letterSpacing: 1,
  },
  primaryButton: {
    backgroundColor: theme.colors.primary,
    paddingVertical: 14,
    paddingHorizontal: 28,
    borderRadius: 14,
    width: "100%",
    marginBottom: 16,
  },
  primaryText: {
    color: "#000",
    fontWeight: "700",
    textAlign: "center",
    letterSpacing: 1,
  },
  secondaryButton: {
    borderWidth: 1,
    borderColor: theme.colors.border,
    paddingVertical: 14,
    paddingHorizontal: 28,
    borderRadius: 14,
    width: "100%",
  },
  secondaryText: {
    color: theme.colors.text,
    fontWeight: "600",
    textAlign: "center",
    letterSpacing: 1,
  },
});
