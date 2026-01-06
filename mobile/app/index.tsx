import { View, Text, Pressable, StyleSheet } from "react-native";
import { useRouter } from "expo-router";
import { useTheme } from "../context/ThemeContext";

export default function Index() {
  const router = useRouter();
  const { colors } = useTheme();
  const dynamicStyles = getStyles(colors);

  return (
    <View style={dynamicStyles.container}>
      <Text style={dynamicStyles.system}>SYSTEM.INITIALIZE(USER)</Text>

      <Text style={dynamicStyles.title}>AARCS-X</Text>
      <Text style={dynamicStyles.subtitle}>AUTH_GATEWAY</Text>

      <Pressable
        style={dynamicStyles.primaryButton}
        onPress={() => router.push("/(auth)/login")}
      >
        <Text style={dynamicStyles.primaryText}>EXECUTE_LOGIN</Text>
      </Pressable>

      <Pressable
        style={dynamicStyles.secondaryButton}
        onPress={() => router.push("/(auth)/signup")}
      >
        <Text style={dynamicStyles.secondaryText}>CREATE_ACCOUNT</Text>
      </Pressable>
    </View>
  );
}

const getStyles = (colors: any) => StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.bg,
    justifyContent: "center",
    alignItems: "center",
    padding: 24,
  },
  system: {
    color: colors.primary,
    letterSpacing: 1,
    marginBottom: 8,
  },
  title: {
    color: colors.text,
    fontSize: 36,
    fontWeight: "800",
    letterSpacing: 2,
  },
  subtitle: {
    color: colors.muted,
    marginBottom: 40,
    letterSpacing: 1,
  },
  primaryButton: {
    backgroundColor: colors.primary,
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
    borderColor: colors.border,
    paddingVertical: 14,
    paddingHorizontal: 28,
    borderRadius: 14,
    width: "100%",
  },
  secondaryText: {
    color: colors.text,
    fontWeight: "600",
    textAlign: "center",
    letterSpacing: 1,
  },
});
