# Doctor Appointment Booking Application - Doctor Availability Module

## Overview

This project implements the **Doctor Availability Module** for a doctor appointment booking system. The system is designed to manage doctor availability by listing and adding time slots, as well as allowing patients to reserve those slots. This module follows the **Traditional Layered Architecture** and manages the core logic of doctor availability.

## Business Requirements Implemented

The **Doctor Availability Module** fulfills the following business requirements:

### 1. **Doctor Availability**
- **List Available Time Slots**: As a doctor, you can list all available time slots.
- **Add New Time Slots**: As a doctor, you can add new available time slots, where each time slot includes:
    - **ID**: GUID
    - **Time**: Date (e.g., 22/02/2023 04:30 PM)
    - **Doctor ID**: GUID
    - **Doctor Name**: String
    - **Is Reserved**: Boolean (false by default)
    - **Cost**: Decimal (e.g., 100.0)

### 2. **Appointment Booking (Partially Implemented)**
- **Reserve a Time Slot**: As a patient, you can reserve an available time slot.

## Architecture

### **Traditional Layered Architecture**
This module uses a traditional layered architecture pattern where the logic is separated into distinct layers:
- **Repository Layer**: Manages in-memory storage of time slots.
- **Service Layer**: Contains business logic for managing time slots (adding, listing, copying).
- **Model Layer**: Contains the definition of the data structures used in the application.

### **Module Structure**
1. **TimeSlotRepositoryImpl**: Manages the time slots in memory using a map. Supports adding, listing, and reserving time slots.
2. **Model**: Defines the `TimeSlot` struct with fields such as `ID`, `Time`, `DoctorName`, `IsReserved`, and `Cost`.
3. **Unit Testing**: Tests the core functionality, including adding time slots, reserving them, and listing all available time slots.

## File Structure
    ```
    doctor_availability/
    │
    ├── model/
    │   └── time_slot.go          # Defines the TimeSlot data model
    │
    ├── repository/
    │   └── time_slot_repository_impl.go  # Implements repository logic (in-memory storage)
    │   └── time_slot_repository_impl_test.go  # Unit tests for the repository
    │
    ├── app/
    │   └── time_slot_repository.go  # Interface for TimeSlot repository
    │
    └── main.go                    # Main entry point for the application
