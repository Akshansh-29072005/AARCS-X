import React, { createContext, useContext, useState, useEffect } from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { theme } from "../constants/theme";

type ThemeType = typeof theme.dark;

interface ThemeContextType {
    isDarkMode: boolean;
    toggleTheme: () => void;
    colors: ThemeType;
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined);

export function ThemeProvider({ children }: { children: React.ReactNode }) {
    const [isDarkMode, setIsDarkMode] = useState(true);

    useEffect(() => {
        // Load persisted theme
        const loadTheme = async () => {
            try {
                const storedTheme = await AsyncStorage.getItem("theme");
                if (storedTheme) {
                    setIsDarkMode(storedTheme === "dark");
                }
            } catch (e) {
                console.error("Failed to load theme", e);
            }
        };
        loadTheme();
    }, []);

    const toggleTheme = async () => {
        try {
            const newMode = !isDarkMode;
            setIsDarkMode(newMode);
            await AsyncStorage.setItem("theme", newMode ? "dark" : "light");
        } catch (e) {
            console.error("Failed to save theme", e);
        }
    };

    const colors = isDarkMode ? theme.dark : theme.light;

    return (
        <ThemeContext.Provider value={{ isDarkMode, toggleTheme, colors }}>
            {children}
        </ThemeContext.Provider>
    );
}

export function useTheme() {
    const context = useContext(ThemeContext);
    if (context === undefined) {
        throw new Error("useTheme must be used within a ThemeProvider");
    }
    return context;
}
