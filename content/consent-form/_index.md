+++
title = "Adult Consent Form"
linktitle = "Consent Form"
layout = "page"
menu = "footer"
+++

This consent form is for anyone 18 years old and over. If the participant is under 18 years old please ask the participant's parent or legal guardian to complete the [Junior Consent Form]({{< ref "/consent-form/junior" >}}).

{{< form name="consent-form" action="/consent-form/success/" >}}
  {{< form/fieldset legend="Participant Information" >}}
    {{< form/text name="first-name" label="First Name" required="yes" autocomplete="section-participant given-name" autofocus="yes" >}}
    {{< form/text name="last-name" label="Last Name" required="yes" autocomplete="section-participant family-name" >}}
    {{< form/email name="email" label="E-mail Address" required="yes" autocomplete="section-participant email" >}}
    {{< form/tel name="phone-number" label="Phone Number" required="yes" autocomplete="section-participant tel" >}}
    {{< form/date name="date-of-birth" label="Date of Birth" placeholder="dd/mm/yyyy" required="yes" autocomplete="section-participant bday" >}}
    {{< form/text name="post-code" label="Postcode" required="yes" autocomplete="section-participant postal-code" >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Medical Details" >}}
    <p>Please detail below any physical, mental or medical conditions that may affect the participant's ability to participate in the activities being provided or <strong>state <q>none</q> if no issues</strong>. (Please discuss with me if you are unsure about this).</p>

    {{< form/textarea name="medical-details" label="Medical Details" required="yes" autocomplete="off" >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Emergency Contact" >}}
    <p>Please provide the name and phone number of someone I can contact if you have an emergency and are unable to contact someone yourself.</p>

    {{< form/text name="emergency-name" label="Emergency Contact Name" required="yes" autocomplete="section-emergency name" >}}
    {{< form/tel name="emergency-tel" label="Emergency Contact Number" required="yes" autocomplete="section-emergency tel" >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Participation Statement" >}}
    <p>Climbing can be mentally and physically demanding, with inherent risks and hazards. Our activities are run and supervised by appropriately qualified and experienced instructors. All necessary safety equipment is provided, this equipment is regularly checked and maintained. Small injuries (e.g., cuts and bruises) are likely, but more serious injuries cannot be ruled out. The instructor(s) will strive to keep you safe. Listen to them, and ask them if you have any questions. You may stop participating at any time. You must bring suitable clothing, food and drink. You must not be intoxicated.</p>
    <p>The <abbr title="British Mountaineering Council">BMC</abbr> recognises that climbing and mountaineering are activities with a danger of personal injury or death. Participants in these activities should be aware of and accept these risks and be responsible for their own actions.</p>

    {{< form/checkbox name="participation-statement" required="yes" >}}
      By ticking this box, you agree that the participant details above are correct and that you understand this Participation Statement.
    {{< /form/checkbox >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="COVID-19 Statement" >}}
    <p>Your instructors will take all necessary precautions to keep you and themselves safe from COVID-19, however we cannot fully guarantee there will be no contamination. Your instructors will be adhering to the social distancing guidelines set out by the government, and the safety recommendations set out by our governing bodies at all times. When activities are taking place outside face coverings will not be required unless social distancing cannot be maintained.<p>

    {{< form/checkbox name="covid-statement" required="yes" >}}
      By ticking this box, you agree not to attend the course if you have symptoms of COVID-19 or have been advised to self-isolate by NHS Track & Trace.
    {{< /form/checkbox >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Electronic Waiver" >}}
    <p>By checking here, you are consenting to the use of your electronic signature in lieu of an original signature on paper. You have the right to request that you sign a paper copy instead. By checking here, you are waiving that right. After consent, you may, upon written request to me, obtain a paper copy of an electronic record. No fee will be charged for such copy and no special hardware or software is required to view it. Your agreement to use an electronic signature with me for any documents will continue until such time as you notify me in writing that you no longer wish to use an electronic signature. There is no penalty for withdrawing your consent. You should always make sure that I have a current email address in order to contact you regarding any changes, if necessary.</p>

    {{< form/checkbox name="electronic-waiver" required="yes" >}}
      I consent to using an electronic signature
    {{< /form/checkbox >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Data Protection Policy" >}}
    <p>To comply with GDPR I am required to obtain your consent to collect and store your personal data. Phill Sparks is the Data Controller. I will store your data securely and will not disclose it to any third party. At any time you can request a copy of your data free of charge. GDPR makes provision for you to request to have your data removed and erased. This is known as the ‘Right to erasure/ right to be forgotten’. When data is held for the establishment, exercise or defence of legal claims I have the right to refuse to erase the data. For this reason our insurers insist that I continue to hold your data for 3 years from the date of your most recent visit. In order to be allowed to use the facilities and services provided by Phill Sparks you are required to consent to our <a href="{{< ref "/policy/privacy" >}}" target="_blank">Privacy Policy</a>.</p>

    {{< form/checkbox name="data-protection-policy" required="yes" >}}
      I agree to the data protection policy
    {{< /form/checkbox >}}
  {{< /form/fieldset >}}

  {{< form/fieldset legend="Declaration" >}}
    <p>I am or over the age of 18 years and I understand the terms of the booking. I have read Phill Sparks’ terms and conditions and fully agree with Phill Sparks' Policies and decisions.</p>

    {{< form/text name="declaration-full-name" label="Print Full Name" required="yes" autocomplete="section-declaration name" >}}
    {{< form/date name="declaration-date" label="Today's Date" placeholder="dd/mm/yyyy" required="yes" >}}
  {{< /form/fieldset >}}

  {{< form/recaptcha >}}
  {{< form/submit label="Submit Consent Form" >}}
{{< /form >}}