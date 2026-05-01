<div align="right">
  <a href="README.pt-BR.md">🇧🇷 Português</a> &nbsp;|&nbsp; <strong>🇺🇸 English</strong>
</div>

# Programming Challenges

> A progressive set of exercises covering programming fundamentals, algorithms, and full-stack development.

---

## Table of Contents

- [About](#about)
- [Evaluator Persona — Ada](#evaluator-persona--ada)
- [Exercises](#exercises)
  - [01 · Votes Relative to Total Voters](#01--votes-relative-to-total-voters)
  - [02 · Bubble Sort](#02--bubble-sort)
  - [03 · Factorial](#03--factorial)
  - [04 · Sum of Multiples of 3 or 5](#04--sum-of-multiples-of-3-or-5)
  - [05 · Vehicle Registration](#05--vehicle-registration)
- [Checkpoint Table](#checkpoint-table)
- [General Evaluation Criteria](#general-evaluation-criteria)

---

## About

This repository contains five programming challenges covering object-oriented fundamentals, classic algorithms, and full-stack web application development. Each exercise is independent and implemented in **Go (Golang)**.

**Suggested directory structure:**

```
desafio-tinnova/
├── ex01-votes/
├── ex02-bubble-sort/
├── ex03-factorial/
├── ex04-multiples/
└── ex05-vehicles/
```

---

## Evaluator Persona — Ada

**Ada** is a senior software engineer with 15 years of experience in back-end development, system architecture, and code review. Her style is direct, technical, and constructive — she doesn't just point out mistakes, she explains the *why* behind each correction and ensures the concept is truly understood before moving on.

Ada operates in three roles:

| Role | What Ada does |
|---|---|
| **Teaching** | Introduces each concept with practical examples and analogies before expecting implementation. Explains the theory behind each exercise. |
| **Evaluation** | Reviews submitted code against the checkpoint table. Marks each item as `APPROVED`, `NEEDS REVISION`, or `REJECTED`. |
| **Mentoring** | After delivery, Ada asks targeted questions to confirm the concept was internalized, not just copied. Example: *"explain why Bubble Sort has O(n²) complexity"*. |

> **Ada's principle:** *"Code that works without understanding is just disguised technical debt."*

---

## Exercises

---

### 01 · Votes Relative to Total Voters

#### Description

Given the following electoral dataset:

```
total voters  = 1000
valid votes   = 800
blank votes   = 150
null votes    = 50
```

Implement a **class** with three methods that calculate:

- The percentage of **valid** votes relative to total voters
- The percentage of **blank** votes relative to total voters
- The percentage of **null** votes relative to total voters

> **Hint:** "relative to total" means the divisor is always the total number of voters.
> Example: `nullPercentage = nullVotes / totalVoters × 100`

Use **object-oriented programming**.

#### Required Concepts

| # | Concept | Why it's needed |
|---|---|---|
| 1 | **Classes and instances** | The exercise requires encapsulating data and methods in a class |
| 2 | **Instance attributes** | Store `totalVoters`, `valid`, `blank`, and `null` as object state |
| 3 | **Instance methods** | Each percentage calculation must be a separate method |
| 4 | **Encapsulation** | Protect data and expose only the necessary interface |
| 5 | **Floating-point arithmetic** | Integer division may truncate — understand `float` vs `int` |
| 6 | **Percentage formula** | Mathematical base: `(part / total) × 100` |
| 7 | **Unit tests** | Verify all three methods against the provided values |

---

### 02 · Bubble Sort

#### Description

Given the following array:

```
v = {5, 3, 2, 4, 7, 1, 0, 6}
```

Implement the **Bubble Sort** algorithm to sort the array in ascending order.

**How Bubble Sort works:**

1. Traverse the array comparing adjacent elements (two at a time)
2. Swap positions if the left element is greater than the right one
3. Repeat the above steps `(n - 1)` times, where `n` is the array length
4. After each pass, the largest remaining element "bubbles up" to its final position and no longer needs to be traversed

**First pass example:**

```
(5  3) 2  4  7  1  0  6   → swap   → 3  5  2  4  7  1  0  6
 3 (5  2) 4  7  1  0  6   → swap   → 3  2  5  4  7  1  0  6
 3  2 (5  4) 7  1  0  6   → swap   → 3  2  4  5  7  1  0  6
 3  2  4 (5  7) 1  0  6   → keep   → 3  2  4  5  7  1  0  6
 3  2  4  5 (7  1) 0  6   → swap   → 3  2  4  5  1  7  0  6
 3  2  4  5  1 (7  0) 6   → swap   → 3  2  4  5  1  0  7  6
 3  2  4  5  1  0 (7  6)  → swap   → 3  2  4  5  1  0  6 [7]
```

`7` is now in its final position. The second pass does not need to include it.

#### Required Concepts

| # | Concept | Why it's needed |
|---|---|---|
| 1 | **Arrays** | The base data structure of this exercise |
| 2 | **Nested loops** | Outer loop (n-1 passes) + inner loop (comparisons) |
| 3 | **Swap operation** | Exchanging two elements requires a temporary variable or destructuring |
| 4 | **Element comparison** | Conditionals to decide whether to swap or keep |
| 5 | **Algorithm complexity** | Understanding O(n²) worst case and why multiple passes are needed |
| 6 | **Incremental optimization** | Reducing the comparison window at each pass |
| 7 | **Array bounds** | Avoiding `IndexOutOfBounds` when comparing adjacent pairs |

---

### 03 · Factorial

#### Description

Implement a program that calculates the **factorial** of any non-negative integer.

**Formal definition:**

```
factorial(n) = 1                       if n = 0
factorial(n) = n × factorial(n - 1)   if n > 0
```

**Expected test cases:**

```
0! = 1
1! = 1
2! = 2
3! = 6
4! = 24
5! = 120
6! = 720
```

> **Note:** `0! = 1` because the empty product (no factors) is defined as 1.

#### Required Concepts

| # | Concept | Why it's needed |
|---|---|---|
| 1 | **Recursion** | The mathematical definition is naturally recursive |
| 2 | **Base case** | `factorial(0) = 1` terminates the recursion; without it, stack overflow occurs |
| 3 | **Recursive case** | `n × factorial(n - 1)` reduces the problem size |
| 4 | **Call stack** | Understanding how each call is pushed and popped from memory |
| 5 | **Pure functions** | Output depends only on the input — no side effects |
| 6 | **Input validation** | Handle negative numbers and non-integers |
| 7 | **Unit tests** | Cover all provided cases, including the `0!` edge case |

---

### 04 · Sum of Multiples of 3 or 5

#### Description

Implement a function that receives a number **X** as a parameter and returns the **sum of all natural numbers below X that are multiples of 3 or 5**.

**Example:**

For `X = 10`, the multiples of 3 or 5 below 10 are: `3, 5, 6, 9`

```
3 + 5 + 6 + 9 = 23
```

> **Note:** "below X" means strictly less than X — X itself is not included.

#### Required Concepts

| # | Concept | Why it's needed |
|---|---|---|
| 1 | **Modulo operator (`%`)** | Check divisibility: `n % 3 === 0` or `n % 5 === 0` |
| 2 | **Loop** | Iterate from 1 to X-1 |
| 3 | **Compound conditional (OR)** | A number is included if it's a multiple of 3 **or** 5 |
| 4 | **Accumulator** | Variable that sums the eligible values |
| 5 | **Function parameters** | The function must be generic, not hardcoded for X=10 |
| 6 | **Edge cases** | X = 0, X = 1, negative X — what should be returned? |
| 7 | **Pure functions** | Same input always produces the same output |

---

### 05 · Vehicle Registration

#### Description

Build a **full-stack** application consisting of:

- **Back-end:** RESTful JSON API written in **Go (Golang)** using all HTTP verbs
- **Front-end:** Single Page Application (SPA) consuming the API

**Data model:**

```
vehicle:     string
brand:       string
year:        integer
description: text
sold:        boolean
created:     datetime
updated:     datetime
```

**Required features:**

- Create, fully update, partially update, and delete vehicles
- Display the count of **unsold** vehicles
- Distribution of vehicles by **decade of manufacture** (e.g., 1990 → 15 vehicles)
- Distribution of vehicles by **manufacturer** (e.g., Ford → 14 vehicles)
- List vehicles registered in the **last week**
- Validate brand consistency (reject misspellings like *Forde*, *Volksvagen*)

**API endpoints:**

| Method | Route | Description |
|---|---|---|
| `GET` | `/veiculos` | Returns all vehicles |
| `GET` | `/veiculos?marca={brand}&ano={year}&cor={color}` | Filter by parameters |
| `GET` | `/veiculos/{id}` | Returns vehicle details |
| `POST` | `/veiculos` | Creates a new vehicle |
| `PUT` | `/veiculos/{id}` | Fully updates a vehicle |
| `PATCH` | `/veiculos/{id}` | Partially updates a vehicle |
| `DELETE` | `/veiculos/{id}` | Deletes a vehicle |

#### Required Concepts

| # | Concept | Why it's needed |
|---|---|---|
| 1 | **REST architecture** | All endpoints follow RESTful conventions |
| 2 | **HTTP verbs** | GET, POST, PUT, PATCH and DELETE with correct semantics |
| 3 | **JSON** | Data exchange format between front-end and back-end |
| 4 | **CRUD** | The four basic operations on the vehicle resource |
| 5 | **Database** | Persistence layer for the data model (relational or NoSQL) |
| 6 | **ORM / Query Builder** | Mapping between objects and database tables |
| 7 | **Data validation** | Invalid brands must be rejected at the input layer |
| 8 | **Filters and query params** | Listing endpoint with optional filters |
| 9 | **Aggregations / GROUP BY** | Distribution by decade and by manufacturer |
| 10 | **Date filtering** | Vehicles registered in the last week (`created >= now - 7 days`) |
| 11 | **SPA (front-end)** | Asynchronous API communication without full page reloads |
| 12 | **Unit tests** | Coverage of endpoints and business rules |
| 13 | **Design Patterns** | e.g., Repository, Service Layer, MVC |
| 14 | **Clean Code** | Expressive names, small functions, no duplication |

---

## Checkpoint Table

Use this table to track progress across all exercises. Ada evaluates each item as `✅ Approved`, `⚠️ Needs Revision`, or `❌ Rejected`.

### Cross-Exercise Checkpoints

| # | Checkpoint | Ex 01 | Ex 02 | Ex 03 | Ex 04 | Ex 05 |
|---|---|:---:|:---:|:---:|:---:|:---:|
| 1 | Code compiles/runs without errors | ☐ | ☐ | ☐ | ☐ | ☐ |
| 2 | Result is mathematically correct | ☐ | ☐ | ☐ | ☐ | ☐ |
| 3 | All provided test cases pass | ☐ | ☐ | ☐ | ☐ | ☐ |
| 4 | Edge cases are handled correctly | ☐ | ☐ | ☐ | ☐ | ☐ |
| 5 | Code uses the concepts required by the exercise | ☐ | ☐ | ☐ | ☐ | ☐ |
| 6 | Variable and function names are expressive | ☐ | ☐ | ☐ | ☐ | ☐ |
| 7 | No unnecessary code duplication | ☐ | ☐ | ☐ | ☐ | ☐ |
| 8 | Unit tests with adequate coverage | ☐ | ☐ | ☐ | ☐ | ☐ |
| 9 | Candidate can explain the solution verbally | ☐ | ☐ | ☐ | ☐ | ☐ |

### Specific Checkpoints — Ex 01 (Votes)

| # | Criterion | Status |
|---|---|:---:|
| 1.1 | A class exists with attributes for the electoral data | ☐ |
| 1.2 | Three separate percentage methods are implemented | ☐ |
| 1.3 | Division uses floating-point (result is not truncated) | ☐ |
| 1.4 | Results for the example values are 80%, 15%, and 5% | ☐ |

### Specific Checkpoints — Ex 02 (Bubble Sort)

| # | Criterion | Status |
|---|---|:---:|
| 2.1 | Algorithm uses two nested loops (or equivalent) | ☐ |
| 2.2 | Swap is performed correctly without losing values | ☐ |
| 2.3 | Comparison window shrinks at each pass | ☐ |
| 2.4 | Array `{5,3,2,4,7,1,0,6}` results in `{0,1,2,3,4,5,6,7}` | ☐ |
| 2.5 | Candidate explains why `n-1` passes are needed | ☐ |

### Specific Checkpoints — Ex 03 (Factorial)

| # | Criterion | Status |
|---|---|:---:|
| 3.1 | Solution uses recursion | ☐ |
| 3.2 | Base case `factorial(0) = 1` is implemented | ☐ |
| 3.3 | `5! = 120` and `6! = 720` are returned correctly | ☐ |
| 3.4 | Invalid inputs (negatives) are handled | ☐ |
| 3.5 | Candidate explains what happens on the call stack | ☐ |

### Specific Checkpoints — Ex 04 (Multiples)

| # | Criterion | Status |
|---|---|:---:|
| 4.1 | Function accepts X as a parameter (not hardcoded) | ☐ |
| 4.2 | Modulo operator is used to check divisibility | ☐ |
| 4.3 | For `X = 10`, the return value is `23` | ☐ |
| 4.4 | Numbers that are multiples of both (e.g., 15) are not double-counted | ☐ |
| 4.5 | Edge cases with X ≤ 0 are handled | ☐ |

### Specific Checkpoints — Ex 05 (Vehicle Registration)

| # | Criterion | Status |
|---|---|:---:|
| 5.1 | API returns JSON responses with correct HTTP status codes | ☐ |
| 5.2 | All 7 endpoints are implemented and functional | ☐ |
| 5.3 | PUT updates all fields; PATCH updates partially | ☐ |
| 5.4 | Count of unsold vehicles is correct | ☐ |
| 5.5 | Grouping by decade of manufacture works correctly | ☐ |
| 5.6 | Grouping by manufacturer works correctly | ☐ |
| 5.7 | Last-week filter works correctly | ☐ |
| 5.8 | Invalid brands are rejected with a clear error message | ☐ |
| 5.9 | Front-end SPA consumes the API without full page reloads | ☐ |
| 5.10 | Unit tests cover endpoints and business rules | ☐ |

---

## General Evaluation Criteria

These criteria apply to all exercises and are evaluated transversally by Ada:

| Criterion | Description |
|---|---|
| **Setup ease** | Project must run with minimal commands; clear instructions in the README |
| **Performance** | Solutions must not have unnecessarily high complexity |
| **Clean code** | Readable, no dead code, no magic numbers, well named |
| **Code documentation** | Comments where intent is not obvious from the code itself |
| **Project documentation** | README with setup, run, and test instructions |
| **Best practices** | Cohesive git commits, `.gitignore`, no credentials in the repository |
| **Design Patterns** | Conscious use of patterns where applicable (especially in Ex 05) |

---

> **Ada's reminder:** *"Before submitting, run all your tests, review your commits, and verify that someone can run the project from scratch following only your README."*
