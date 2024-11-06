import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
    return {
        features: [
            {
                title: "Customizable",
                description:
                    "Tailor the template to your specific needs and preferences.",
            },
            {
                title: "Fast",
                description:
                    "Optimized for speed to ensure your website loads quickly.",
            },
            {
                title: "Responsive",
                description: "Looks great on desktop, tablet, and mobile devices.",
            },
            {
                title: "Accessible",
                description:
                    "Designed with accessibility in mind to reach all users.",
            },
        ],
        pricingPlans: [
            {
                name: "Free",
                price: "$0",
                description: "For small side projects.",
                features: ["1 user", "Basic support", "1 project"],
            },
            {
                name: "Pro",
                price: "$15",
                description: "For professional developers.",
                features: ["5 users", "Priority support", "10 projects"],
            },
        ],
        faqs: [
            {
                question: "Is this template free to use?",
                answer: "Yes, this template is free to use for both personal and commercial projects.",
            },
            {
                question: "Can I use this template for client projects?",
                answer: "You're free to use this template for client projects.",
            },
            {
                question: "How do I customize this template?",
                answer: "You can customize this template by modifying the Svelte components and Tailwind CSS classes.",
            },
            {
                question: "Is there a premium version available?",
                answer: "Currently, there isn't a premium version. This is the full template available for free.",
            },
        ],
    };
};