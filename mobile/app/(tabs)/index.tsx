import { View, Text } from "react-native";
import { useEffect, useState } from "react";
import { getStudents } from "../../lib/api";

export default function HomeScreen() {
  const [count, setCount] = useState<number | null>(null);

  useEffect(() => {
    getStudents()
      .then((data) => setCount(data.length))
      .catch((err) => console.error(err));
  }, []);

  return (
    <View style={{ flex: 1, alignItems: "center", justifyContent: "center" }}>
      <Text style={{ fontSize: 22, color : "cyan" }}>AARCS-X Mobile</Text>
      <Text style={{ marginTop: 10 , color: 'gray' , fontSize: 16 }}>
        Students count: {count ?? "loading..."}
      </Text>
    </View>
  );
}
