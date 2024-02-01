package main

import (
    "fmt"
    "time"

    "github.com/tebeka/selenium"
    "github.com/tebeka/selenium/chrome"
)

// func main() {
    // Start Selenium WebDriver
    caps := selenium.Capabilities{
        "browserName": "chrome",
    }
    chromeCaps := chrome.Capabilities{}
    caps.AddChrome(chromeCaps)
    wd, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
    if err != nil {
        fmt.Println("Failed to open session:", err)
        return
    }
    defer wd.Quit()

    // Navigate to the AWS Management Console
    if err := wd.Get("https://aws.amazon.com/console/"); err != nil {
        fmt.Println("Failed to load AWS Management Console:", err)
        return
    }

    // Wait for the page to load
    time.Sleep(20 * time.Second)

    // Find the "Sign In to the Console" button and click it
    signInButton, err := wd.FindElement(selenium.ByClassName, "lb-btn-p-primary")
    if err != nil {
        fmt.Println("Failed to find sign-in button:", err)
        return
    }
    if err := signInButton.Click(); err != nil {
        fmt.Println("Failed to click sign-in button:", err)
        return
    }

    // Wait for the sign-in page to load
    time.Sleep(20 * time.Second)

    // You can continue with further actions on the sign-in page if needed

    
    // Find the "Sign in to an existing AWS account" button and click it
    signInButtonn, errr := wd.FindElement(selenium.ByClassName, "awsui-button-variant-normal")
    if err != nil {
        fmt.Println("Failed to find sign-in button:", err)
        return
    }
    if err := signInButtonn.Click(); errr != nil {
        fmt.Println("Failed to click sign-in button:", err)
        return
    }

    // Wait for the next page to load
    time.Sleep(60 * time.Second)



    // Find the text box for the Root user email address and enter the email
    emailTextBox, err := wd.FindElement(selenium.ByID, "resolving_input")
    if err != nil {
        fmt.Println("Failed to find email text box:", err)
        return
    }
    if err := emailTextBox.SendKeys("Koushik.Anchuri2001@gmail.com"); err != nil {
        fmt.Println("Failed to enter email:", err)
        return
    }

    // Find and click the Next button
    nextButton, err := wd.FindElement(selenium.ByID, "next_button")
    if err != nil {
        fmt.Println("Failed to find Next button:", err)
        return
    }
    if err := nextButton.Click(); err != nil {
        fmt.Println("Failed to click Next button:", err)
        return
    }

    // Wait for the next page to load
    time.Sleep(30 * time.Second)


    // Find the Password text box and enter the password
    passwordTextBox, err := wd.FindElement(selenium.ByID, "password")
    if err != nil {
        fmt.Println("Failed to find password text box:", err)
        return
    }
    if err := passwordTextBox.SendKeys("Anchuri@87509"); err != nil {
        fmt.Println("Failed to enter password:", err)
        return
    }

    // Find and click the Sign In button
    signInButtonnn, errrr := wd.FindElement(selenium.ByID, "signin_button")
    if err != nil {
        fmt.Println("Failed to find Sign In button:", err)
        return
    }
    if err := signInButtonnn.Click(); errrr != nil {
        fmt.Println("Failed to click Sign In button:", err)
        return
    }

    // Wait for the next page to load after signing in
    time.Sleep(20 * time.Second)

    // Continue with further actions on the signed-in page if needed

}
