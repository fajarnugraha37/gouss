<script>
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";

    let url = "";
    let alias = "";
    let shortenedUrl = "";

    const shortenUrl = async () => {
        try {
            const response = await fetch("/api/shorten", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ url, alias }),
            });
            const data = await response.json();
            shortenedUrl = data.shortenedUrl;
        } catch (error) {
            console.error("Error shortening URL:", error);
        }
    };
</script>

<main class="container mx-auto p-4 text-white">
    <section class="hero text-center mb-8">
        <h1 class="text-5xl font-bold mb-2">GOUSS URL Shortener</h1>
        <p class="text-xl">Fast, Reliable, and Secure URL Management</p>
    </section>
    <section class="form-section bg-white text-golangBlue p-6 rounded-md shadow-md">
        <div class="mb-4">
            <label for="url" class="block text-sm font-medium mb-2">Enter your URL</label>
            <Input
                id="url"
                type="url"
                bind:value={url}
                class="mt-1 block w-full p-2 border border-gray-300 rounded-md text-gray-900"
                placeholder="https://example.com"
            />
        </div>
        <div class="mb-4">
            <label for="alias" class="block text-sm font-medium mb-2"
                >Custom Alias (Optional)</label
            >
            <Input
                id="alias"
                type="text"
                bind:value={alias}
                class="mt-1 block w-full p-2 border border-gray-300 rounded-md text-gray-900"
                placeholder="Custom alias"
            />
        </div>
        <Button on:click={shortenUrl} class="mt-4 bg-golangBlue text-white p-2 rounded-md">Shorten URL</Button>
        {#if shortenedUrl}
            <div class="mt-4 p-4 border border-green-500 rounded-md">
                <p class="text-green-600">
                    Your shortened URL: <a
                        href={shortenedUrl}
                        class="text-blue-500">{shortenedUrl}</a
                    >
                </p>
            </div>
        {/if}
    </section>
</main>

<style>
    .container {
        max-width: 100%;
        width: 100%;
        padding: 1rem;
    }
    @media (min-width: 640px) {
        .container {
            max-width: 640px;
        }
    }
    @media (min-width: 768px) {
        .container {
            max-width: 768px;
        }
    }
    @media (min-width: 1024px) {
        .container {
            max-width: 1024px;
        }
    }
    .hero {
        background: linear-gradient(135deg, #00add8 0%, #007c8e 100%);
        padding: 4rem 1rem;
        border-radius: 0.5rem;
    }
    .form-section {
        background: #fff;
    }
    label {
        color: #007c8e;
    }
</style>
