#!/bin/bash

BASE_URL="http://localhost:8000/api/v1"
TIMESTAMP=$(date +%s)

# Helper function to register and login a user
register_login_user() {
    local suffix=$1
    local num=$2
    local email="testuser_${suffix}_${TIMESTAMP}@example.com"
    local name="Test User ${suffix} ${TIMESTAMP}"
    local phone="+91987654320${num}"
    local password="password123"

    echo "Registering user $suffix ($email)..."
    REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
      -H "Content-Type: application/json" \
      -d "{
        \"name\": \"$name\",
        \"email\": \"$email\",
        \"phone_number\": \"$phone\",
        \"password\": \"$password\"
      }")
    
    echo "Register Response ($suffix): $REGISTER_RESPONSE"
    TOKEN=$(echo $REGISTER_RESPONSE | jq -r '.token')
    
    if [ "$TOKEN" == "null" ] || [ -z "$TOKEN" ]; then
        echo "Error: Failed to get token during registration ($suffix). Response: $REGISTER_RESPONSE"
        # exit 1
    fi
    
    # Get User ID
    echo "Getting User ID for $suffix..."
    ME_RESPONSE=$(curl -s -X GET "$BASE_URL/auth/protected/me" \
      -H "Authorization: Bearer $TOKEN")
    echo "Me Response ($suffix): $ME_RESPONSE"
    USER_ID=$(echo $ME_RESPONSE | jq -r '.user_id')
    
    echo "$TOKEN|$USER_ID"
}

echo "--- Starting Smoke Test ---"

# 1. Register/Login User A (Owner)
echo "Step 1: Registering User A (Owner)..."
USER_A_DATA=$(register_login_user "A" "1")
TOKEN_A=$(echo $USER_A_DATA | cut -d'|' -f1)
ID_A=$(echo $USER_A_DATA | cut -d'|' -f2)
echo "User A ID: $ID_A"

# 2. Create Institute using User A
echo "Step 2: Creating Institute using User A..."
INSTITUTE_NAME="Test Institute ${TIMESTAMP}"
INSTITUTE_CODE="TI_${TIMESTAMP}"
INSTITUTE_EMAIL="inst_${TIMESTAMP}@example.com"

CREATE_INST_RESPONSE=$(curl -s -X POST "$BASE_URL/institutions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN_A" \
  -d "{
    \"name\": \"$INSTITUTE_NAME\",
    \"code\": \"$INSTITUTE_CODE\",
    \"official_email\": \"$INSTITUTE_EMAIL\",
    \"address\": \"123 Test St\",
    \"district\": \"Test District\",
    \"state\": \"Test State\",
    \"country\": \"Test Country\"
  }")

echo "Create Institute Response: $CREATE_INST_RESPONSE"
INST_ID=$(echo $CREATE_INST_RESPONSE | jq -r '.id')
if [ "$INST_ID" == "null" ] || [ -z "$INST_ID" ]; then
    echo "FAILURE: Failed to create institute. Response: $CREATE_INST_RESPONSE"
else
    echo "Institute Created with ID: $INST_ID"
fi

# 3. Register/Login User B (Potential Admin)
echo "Step 3: Registering User B (Target Admin)..."
USER_B_DATA=$(register_login_user "B" "2")
TOKEN_B=$(echo $USER_B_DATA | cut -d'|' -f1)
ID_B=$(echo $USER_B_DATA | cut -d'|' -f2)
echo "User B ID: $ID_B"

# 4. User A promotes User B to Admin
echo "Step 4: User A promoting User B to Admin..."
PROMOTE_RESPONSE=$(curl -s -X POST "$BASE_URL/institutions/make-admin" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN_A" \
  -d "{
    \"institution_id\": $INST_ID,
    \"user_id\": $ID_B
  }")
echo "Promote Response: $PROMOTE_RESPONSE"

# 5. Verify User A (Owner) can see their institutions
echo "Step 5: Verifying User A can see their institutions..."
LIST_INST_RESPONSE=$(curl -s -X GET "$BASE_URL/institutions" \
  -H "Authorization: Bearer $TOKEN_A")
echo "List Response (User A): $LIST_INST_RESPONSE"

# 6. Verify User B (Admin) cannot make admins
echo "Step 6: Registering User C and testing if User B can make them admin..."
USER_C_DATA=$(register_login_user "C" "3")
ID_C=$(echo $USER_C_DATA | cut -d'|' -f2)

PROMOTE_FAIL_RESPONSE=$(curl -s -X POST "$BASE_URL/institutions/make-admin" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN_B" \
  -d "{
    \"institution_id\": $INST_ID,
    \"user_id\": $ID_C
  }")
echo "Promote Fail Response (User B): $PROMOTE_FAIL_RESPONSE"

echo "--- Smoke Test Completed ---"
