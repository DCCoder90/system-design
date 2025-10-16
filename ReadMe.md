# System Design

This repository is a collection of system design exercises I'm undertaking to practice and improve my skills. While some projects may include proof-of-concept code, the primary focus is on the design process, architecture, and iterating on complex systems. This is not intended to be a repository of production-ready products.

## How to Use This Repo

Each project is self-contained within its own directory and follows an iterative structure.

  * Project Hub: Every project folder (e.g., `/Notification System`) contains a main `ReadMe.md` file. This file serves as the central hub, outlining the project's overall goals and linking to the different design versions.
  * Iterative Versions: Inside each project, you will find versioned directories (`/v1`, `/v2`, etc.). Each version represents a different stage of the design process, starting with a MVP and progressively adding features and complexity.

### Typical Structure

```
/ProjectName
├── ReadMe.md           <-- Main project overview and links to versions
├── /v1                 <-- Version 1: MVP/Proof of Concept
│   ├── ReadMe.md       <-- Detailed write-up for v1 architecture & scope
└── /v2                 <-- Version 2: Iteration with new features
    ├── ReadMe.md       <-- Detailed write-up for v2 architecture & scope
```

## Projects

### Notification System

The [Notification System Project](./Notification%20System/ReadMe.md) is an exercise to design the architecture and core implementation of a highly scalable notification system. This system will be designed to reliably deliver billions of messages across multiple channels, including features such as scheduling, templates, and idempotency.

  * **Status:** In Progress.

## Contributions

As this is a personal repository for learning and practice, I am not actively seeking contributions. However, if you have suggestions or feedback on any of the designs, feel free to open an issue!