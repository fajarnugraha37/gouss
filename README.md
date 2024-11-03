<h1 align="center">
  <a href="https://github.com/fajarnugraha37/gouss">
    <picture>
      <img alt="GOUSS" src="https://raw.githubusercontent.com/fajarnugraha37/gouss/refs/heads/main/docs/logo-4.png">
    </picture>
  </a>
</h1>
<p align="center">
    <em><b>GOUSS</b> A simple, reliable URL shortener featuring a cute, friendly mascot. Designed for ease of use, secure URL management, analytics, and integration options. Perfect for developers seeking a streamlined link management tool with personality.
    </em>
</p>

---
## ⚠️ **Attention**

This project, GOUSS, is currently under active development and is not yet ready for use.

## Features

- [] **Basic URL Shortening**: Generate short, unique URLs for long links.
- [] **Persistent Storage**: Uses a database for URL storage to ensure data persistence.
- [] **Admin Dashboard**: managing links and viewing statistics.
- [] **User Authentication**: Support for user accounts and secure access.
- [] **OAuth Support**: Integrates with GitHub, Google, Facebook, etc.
- [] **Subscription Models**: Monetized plans.
- [] **Advanced Analytics**: Track click counts, referrer data, device type, etc.
- [] **Extensive API**: RESTful API for integration.

## Roadmap

This project will evolve from a basic prototype to a real product that can be used in production with the following roadmap:
### Phase 1:
- [] **Core URL Shortening**: Implement the basic URL shortening mechanism.
- [] **In-Memory Storage**: Temporarily store shortened URLs in memory.
- [] **Simple Redirect**: Handle HTTP redirects from shortened URLs to the original URLs.
- [] **Basic API Documentation**: Provide endpoints for creating and accessing shortened URLs.
- [] **HTTPS**: HTTPS support

### Phase 2:
- [] **Persistent Database**: Store URLs in a persistent database (PostgreSQL).
- [] **URL Validation**: Ensure input URLs are valid and secure.
- [] **Rate Limiting**: Prevent abuse with request rate limiting.
- [] **Error Handling**: Standardize error messages and status codes.
- [] **Unit Testing**: Implement unit tests for core functions.

### Phase 3:
- [] **Caching URLs**: Use Redis for frequent URLs.
- [] **Expiration Dates for URLs**: Allow users to set expiration dates for URLs.
- [] **URL Statistics**: Track and display click counts for each shortened URL.
- [] **API Enhancements**: Add endpoints for retrieving URL statistics and managing expiration settings.

### Phase 4:
- [] **Custom Short URLs**: Enable custom aliases for URLs.
- [] **Containerization**: Dockerize the application for easier deployment.
- [] **Logging & Monitoring**: Integrate with logging and monitoring solutions.
- [] **User Accounts & Authentication**: Implement user registration and login with JWT.
- [] **Notification**: Send notifications for specific actions like URLC creation or deletion.

### Phase 5:
- [] **QR Code Generation**: Automatically generate QR codes for shortened URLs.
- [] **OAuth Support**: Integrate with OAuth providers (e.g., GitHub, Google, Okta, etc) for authentication.
- [] **Advanced Analytics**: Track detailed statistics, including referrer and device type.
- [] **A/B Testing**: Allow users to test variations of URLs in campaigns.
- [] **Collaboration & Team Management**: Allow team-based link management.

### Phase 6: Payment Integrations and Subscription Models
- [] **Feature Access Control**: Control access to certain features based on the user's subscription plan.
- [] **Subscription Plans**: Offer various plans, such as Free, Basic, and Pro, with different limits and features.
- [] **Payment Integration**: Add support for payment gateways (e.g., Stripe, PayPal) to enable transactions.
- [] **Billing Management**: Enable users to view and manage their billing details.
- [] **Trial Periods and Discounts**: Allow users to try premium features for a limited time and offer discounts.

## Demo

## Architecture

## Quick Start

### Prerequisites
- Go 1.22.3+
- Node v18.17.0+
- NPM 9.6.7+

### Installation

### Tests

### Benchmark
