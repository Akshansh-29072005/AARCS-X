#!/bin/bash

BASE_URL="http://localhost:8000/api/v1"
TIMESTAMP=$(date +%s)
EMAIL="testuser_${TIMESTAMP}@example.com"
NAME="Test User ${TIMESTAMP}"
PHONE="+1234567890"
PASSWORD="password123"

echo "Using Email: $EMAIL"

# 1. Register User
echo "Step 1: Registering user..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"$NAME\",
    \"email\": \"$EMAIL\",
    \"phone_number\": \"$PHONE\",
    \"password\": \"$PASSWORD\"
  }")

echo "Register Response: $REGISTER_RESPONSE"

# 2. Login
echo "Step 2: Logging in..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }")

echo "Login Response: $LOGIN_RESPONSE"
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')

if [ "$TOKEN" == "null" ] || [ -z "$TOKEN" ]; then
  echo "Error: Failed to get token."
  exit 1
fi

echo "Token obtained: ${TOKEN:0:20}..."

# 3. Create Institute
echo "Step 3: Creating institute..."
INSTITUTE_NAME="Test Institute ${TIMESTAMP}"
INSTITUTE_CODE="TI_${TIMESTAMP}"
INSTITUTE_EMAIL="inst_${TIMESTAMP}@example.com"

CREATE_INST_RESPONSE=$(curl -s -X POST "$BASE_URL/institutions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
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
  echo "Error: Failed to create institute."
  # exit 1 # Don't exit, let's see why
fi

# 4. Read Institutes (List)
echo "Step 4: Listing institutes..."
LIST_INST_RESPONSE=$(curl -s -X GET "$BASE_URL/institutions" \
  -H "Authorization: Bearer $TOKEN")

echo "List Institutes Response: $LIST_INST_RESPONSE"

# 5. Read Institute by ID
if [ "$INST_ID" != "null" ] && [ -n "$INST_ID" ]; then
  echo "Step 5: Getting institute by ID ($INST_ID)..."
  GET_INST_RESPONSE=$(curl -s -X GET "$BASE_URL/institutions/$INST_ID" \
    -H "Authorization: Bearer $TOKEN")
  echo "Get Institute By ID Response: $GET_INST_RESPONSE"
fi
